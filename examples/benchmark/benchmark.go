package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/cjongseok/mtproto"
	"github.com/cjongseok/ntped"
	"github.com/cjongseok/slog"
	"github.com/montanaflynn/stats"
)

type condition func(float64) bool

type updateBenchmarker struct {
	latencies   []int64
	updateNum   int
	updateCount int
	join        sync.WaitGroup
}

func usage() {
	fmt.Println("USAGE: ./benchmark <APIID> <APIHASH> <PHONE> <ADDR> <PING> <BENCH> [ARGS]")
	fmt.Println("")
	fmt.Println("    APIID 	  means Telegram API id. If you do not have it yet, make it at ")
	fmt.Println("	    		  https://my.telegram.org/apps")
	fmt.Println("    APIHASH   means hashcode of <APIID>. It is published together with API id.")
	fmt.Println("    PHONE 	  means already authenticated phonenumber. If you have never")
	fmt.Println("      		  signed in to Telegram with simple shell, authenticate your")
	fmt.Println("     		  phone through simpleshell first.")
	fmt.Println("    ADDR 	  means Telegram server address, e.g., 91.108.56.165:443")
	fmt.Println("    PING 	  means ping interval in milli seconds.")
	fmt.Println("")
	fmt.Println("BENCHES:")
	fmt.Println("  UpdateLatency <#UPDATES>")
	fmt.Println("  UpdateLatencyWithPolling <#UPDATES> <CHANID> <CHANHASH> <INTERVAL>")
	fmt.Println("    #UPDATES  means number of updates to subscribe from the server. Setting")
	fmt.Println("     		  it as 0 subscribes updates forever.")
	fmt.Println("    CHANID 	  means channel id.")
	fmt.Println("    CHANHASH  means channel hash.")
	fmt.Println("    #UPDATES  means number of updates to subscribe from the server. Setting")
	fmt.Println("     		  it as 0 subscribes updates forever.")
	fmt.Println("    INTERVAL	  means polling interval in milli seconds.")
}

func main() {
	const (
		minArgsNum = 7
	)
	if len(os.Args) < minArgsNum {
		usage()
		return
	}

	apiid64, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		usage()
		return
	}
	apiid := int32(apiid64)
	apihash := os.Args[2]
	phoneNumber := os.Args[3]
	target := os.Args[4]
	intervalInMs := os.Args[5]
	bench := os.Args[6]
	var args []string

	benchtime := time.Now()
	y, mon, d := benchtime.Date()
	h, min, s := benchtime.Clock()
	dateString := fmt.Sprintf("%02d%02d%02d-%02d%02d%02d", (y % 100), mon, d, h, min, s)
	lfilename := fmt.Sprintf("%s_%s_benchmark.log", dateString, target)
	bfilename := fmt.Sprintf("%s_%s_benchmark.bch", dateString, target)

	// Set logfile
	lfile, err := os.OpenFile(lfilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("error opening file: ", err)
	}
	defer lfile.Close()
	log.SetOutput(lfile)
	//log.SetOutput(os.Stdout)

	// Set slog.Benchfile
	bfile, err := os.OpenFile(bfilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("error opening file: ", err)
	}
	defer bfile.Close()
	//bfile := os.Stdout
	slog.SetBenchOutput(bfile)

	slog.Benchln(main, "==== Settings ====")
	slog.Benchln(main, "target: ", target)
	slog.Benchf(main, "ping: %s ms\n", intervalInMs)

	var intervalDuration time.Duration
	interval64, err := strconv.ParseInt(intervalInMs, 10, 0)
	if err == nil {
		intervalDuration = time.Duration(interval64 * int64(time.Millisecond))
	}
	config := benchConfig{apiid, apihash, phoneNumber, target, intervalDuration}

	switch bench {
	case "UpdateLatency":
		if len(os.Args) != 1+minArgsNum {
			usage()
			return
		}
		args = os.Args[minArgsNum:]
		updateNum, err := strconv.Atoi(args[0])
		handleArgError(err)
		slog.Benchln(main, "updateNum: ", updateNum)
		slog.Benchln(main, "==================")
		benchmarkUpdateLatency(config, updateNum)

	case "UpdateLatencyWithPolling":
		if len(os.Args) != 4+minArgsNum {
			usage()
			return
		}
		args = os.Args[minArgsNum:]
		chanid, err := strconv.ParseInt(args[0], 10, 32)
		handleArgError(err)
		chanhash, err := strconv.ParseInt(args[1], 10, 64)
		handleArgError(err)
		updateNum, err := strconv.Atoi(args[2])
		handleArgError(err)
		interval, err := strconv.ParseInt(args[3], 10, 64)
		handleArgError(err)
		slog.Benchln(main, "updateNum: ", updateNum)
		slog.Benchln(main, "chanid: ", chanid)
		slog.Benchln(main, "chanhash: ", chanhash)
		slog.Benchln(main, "interval: ", interval)
		slog.Benchln(main, "==================")
		intervalInDuration := time.Duration(interval) * time.Millisecond
		benchmarkUpdateLatencyWithPolling(config, updateNum, int32(chanid), int64(chanhash), intervalInDuration)

	default:
	}
}

type benchConfig struct {
	apiid         int32
	apihash       string
	phoneNumber   string
	preferredAddr string
	pingInterval  time.Duration
}
type abort func()
type onMeasuring func(*mtproto.MConn) abort

// Channel polling == get full channel on init poll +
func benchmarkUpdateLatencyWithPolling(config benchConfig, updateNum int, chanid int32, chanhash int64, interval time.Duration) error {
	polling := func(mconn *mtproto.MConn) abort {
		stopPolling := make(chan interface{})
		go func(stop chan interface{}) {
			// get full channel
			var readInboxMaxId int32
			inputchannel := mtproto.TL_inputChannel{chanid, chanhash}
			resp, err := mconn.InvokeBlocked(mtproto.TL_channels_getFullChannel{inputchannel})
			if err != nil {
				slog.Logln(mconn, "Failed to get full channel:", err)
			} else {
				slog.Logln(mconn, "FullChannel:", slog.Stringify(resp))
				readInboxMaxId = (*resp).(mtproto.TL_messages_chatFull).Full_chat.(mtproto.TL_channelFull).Read_inbox_max_id
			}
			slog.Benchln(benchmarkUpdateLatencyWithPolling, "FullChannel.read_inbox_max_id=", readInboxMaxId)

			// start polling
			slog.Logln(mconn, "Start of polling")
			for {
				select {
				case <-stop:
					slog.Logln(mconn, "stop polling")
					return
				case <-time.After(interval):
					slog.Logf(mconn, "poll to %d by %d\n", chanid, chanhash)
					peer := mtproto.TL_inputPeerChannel{chanid, chanhash}
					//resp, err := mconn.MessagesGetHistory(peer, 0, 0, 0, 0, 0, readInboxMaxId + 1)
					//TODO: deadlocked on Session() ?
					resp, err := mconn.MessagesGetHistory(peer, 0, 0, 0, 0, 0, readInboxMaxId)
					if err != nil {
						slog.Logln(mconn, "Failed to poll:", err)
						slog.Benchln(benchmarkUpdateLatencyWithPolling, "Polling Failure:", err)
					} else {
						msgs := (*resp).(mtproto.TL_messages_channelMessages).Unstrip().(mtproto.US_messages_channelMessages).Messages
						slog.Logln(mconn, "Polled history:", slog.Stringify(msgs))
						if len(msgs) > 0 {
							readInboxMaxId = msgs[0].Id
						}
						for _, msg := range msgs {
							slog.Benchf(benchmarkUpdateLatencyWithPolling, "[Polled history %d] %s\n", msg.Id, msg.Message)
						}

					}
				}
			}
			slog.Logln(mconn, "End of polling")
		}(stopPolling)
		abortPolling := func() {
			// stop polling
			close(stopPolling)
		}
		return abortPolling
	}
	return measureUpdateLatency(config, updateNum, polling)
}

func benchmarkUpdateLatency(config benchConfig, updateNum int) error {
	return measureUpdateLatency(config, updateNum, nil)
}

func measureUpdateLatency(config benchConfig, updateNum int, taskOnMeasuring onMeasuring) error {
	slog.EnableLogging()
	mm, mconn, err := initBenchmark(config)
	if err != nil {
		return err
	}

	//slog.DisableLogging()
	slog.Benchln(measureUpdateLatency, "start")
	b := NewUpdateBenchmarker(updateNum)
	mconn.AddUpdateCallback(b)
	var abortTask abort
	if taskOnMeasuring != nil {
		abortTask = taskOnMeasuring(mconn)
	}

	// join the measuring
	b.join.Wait()
	slog.Benchln(measureUpdateLatency, "done")
	slog.EnableLogging()

	// report analytics
	//floats := Float64s(b.latencies)
	floats := stats.LoadRawData(b.latencies)
	rmean, _ := stats.Mean(floats)
	rmedian, _ := stats.Median(floats)
	rmin, _ := stats.Min(floats)
	rmax, _ := stats.Max(floats)
	rvariance, _ := stats.Variance(floats)
	rstd, _ := stats.StandardDeviation(floats)

	filtered := filter(floats, func(x float64) bool { return x > 500 })
	fmean, _ := stats.Mean(floats)
	fmedian, _ := stats.Median(floats)
	fmin, _ := stats.Min(floats)
	fmax, _ := stats.Max(floats)
	fvariance, _ := stats.Variance(floats)
	fstd, _ := stats.StandardDeviation(floats)

	//CAVEAT: it measures latencies by comparing timestamps (unit: sec) embedded in updates and when update callback triggers by NTP clock (unit: ms).
	//You need to keep in mind two things.
	// 1. By the unit difference, the timestamped time \\in [0,1) + timestamp
	//   So remove such noise before analytics.
	//	 e.g., pick latencies over 0.5, which is an expectation when noise uniformly distributes.
	// 2. By the time source difference, if Telegram server is not synced with NTP pool, latencies can incorrect.
	//   So measure latencies from a baseline experiment environment first, and compare it with other environments.
	slog.Benchln(measureUpdateLatency, "====== RESULTS ======")
	slog.Benchf(measureUpdateLatency, "raw latencies: %s\n", slog.Stringify(b.latencies))
	slog.Benchf(measureUpdateLatency, "raw mean:%f, median:%f, min:%f, max:%f, var:%f, std:%f\n", rmean, rmedian, rmin, rmax, rvariance, rstd)
	slog.Benchf(measureUpdateLatency, "filtered latencies: %s\n", slog.Stringify(filtered))
	slog.Benchf(measureUpdateLatency, "filtered mean:%f, median:%f, min:%f, max:%f, var:%f, std:%f\n", fmean, fmedian, fmin, fmax, fvariance, fstd)
	slog.Benchln(measureUpdateLatency, "=====================")

	// Finish the benchmark
	if taskOnMeasuring != nil && abortTask != nil {
		abortTask()
	}
	mm.Finish()
	return nil
}

func filter(ints []float64, cond condition) []float64 {
	size := 0
	for _, x := range ints {
		if cond(x) {
			size++
		}
	}
	filtered := make([]float64, size)
	i := 0
	for _, x := range ints {
		if cond(x) {
			filtered[i] = x
			i++
		}
	}
	return filtered
}

func NewUpdateBenchmarker(updateNum int) *updateBenchmarker {
	b := new(updateBenchmarker)
	b.updateCount = 0
	b.updateNum = updateNum
	b.join = sync.WaitGroup{}
	b.join.Add(1)
	return b
}

func (b *updateBenchmarker) OnUpdate(u mtproto.MUpdate) {
	tic := ntped.UnixMilli()
	updateDate := int64(u.UpdateDate()) * 1e3 // update date in milli
	diff := tic - updateDate
	if updateDate != 0 {
		slog.Benchf(b, "[tdiff = %d ms] got %s\n", diff, slog.Stringify(u))
		b.latencies = append(b.latencies, diff)
		b.updateCount++
		if b.updateNum == 0 {
			// never stop
		} else if b.updateCount >= b.updateNum {
			b.join.Done()
		}
	} else {
		slog.Benchf(b, "got %s\n", slog.Stringify(u))
	}
}

func initBenchmark(config benchConfig) (*mtproto.MManager, *mtproto.MConn, error) {
	const (
		maxNtpRetry = 0
		timeoutInMs = 10000

		appVersion      = "0.0.1"
		deviceModel     = ""
		systemVersion   = ""
		language        = ""
		sessionFileHome = ""
	)

	// sync ntped and sign in Telegram
	join := sync.WaitGroup{}
	results := make(chan interface{})

	// Sync ntped routine
	join.Add(1)
	go func(out chan<- interface{}) {
		defer join.Done()
		err := ntped.Sync(maxNtpRetry, timeoutInMs)
		if err != nil {
			out <- err
		} else {
			out <- true
		}
	}(results)

	// Sign-in routine
	join.Add(1)
	go func(out chan<- interface{}) {
		defer join.Done()
		configuration, err := mtproto.NewConfiguration(config.apiid, config.apihash, appVersion, deviceModel, systemVersion, language, sessionFileHome, config.pingInterval, 0)
		handleError(err)
		manager, err := mtproto.NewManager(configuration)
		handleError(err)

		if !manager.IsAuthenticated(config.phoneNumber) {
			out <- fmt.Errorf("auth does NOT exist")
			return
		}
		mconn, err := manager.LoadAuthentication(config.phoneNumber, config.preferredAddr)
		if err != nil {
			out <- err
		} else {
			out <- manager
			out <- mconn
		}
	}(results)

	// Waiter for results
	go func() {
		join.Wait()
		close(results)
	}()

	// gather results
	waitForResults := true
	ntpedSynced := false
	var mm *mtproto.MManager
	var mconn *mtproto.MConn
	for waitForResults {
		select {
		case result, ok := <-results:
			if !ok {
				waitForResults = false
			}
			switch result.(type) {
			case bool:
				ntpedSynced = true
				slog.Logln(initBenchmark, "clock is synced")
				slog.SetBenchClock(ntped.Now)
			case *mtproto.MManager:
				mm = result.(*mtproto.MManager)
				slog.Logln(initBenchmark, "got mmanager, ", slog.Stringify(mm))
			case *mtproto.MConn:
				mconn = result.(*mtproto.MConn)
				slog.Logln(initBenchmark, "got mconn, ", slog.Stringify(mconn))
			case error:
				slog.Logln(initBenchmark, "got error, ", result.(error))
				return nil, nil, result.(error)
			default:
				slog.Logln(initBenchmark, "Unknown result, ", slog.Stringify(result))
			}
		}
	}
	if mm == nil {
		return nil, nil, fmt.Errorf("Sign-in Failure: no mmanager")
	}
	if mconn == nil {
		return nil, nil, fmt.Errorf("Sign-in Failure: no mconn")
	}
	if !ntpedSynced {
		return nil, nil, fmt.Errorf("NTPed clock Initialization Failure")
	}
	return mm, mconn, nil
}

func handleArgError(err error) {
	if err != nil {
		usage()
		slog.Fatalln(err)
	}
}
func handleError(err error) {
	if err != nil {
		slog.Fatalln(err)
	}
}
