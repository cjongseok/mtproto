package main

import (
	"log"
	"os"
	"fmt"
	"sync"
	"time"
	"strconv"

	"github.com/cjongseok/mtproto"
	"github.com/cjongseok/ntped"
	"github.com/cjongseok/slog"
	"github.com/montanaflynn/stats"
)


func usage() {
	fmt.Println("USAGE: ./benchmark <APIID> <APIHASH> <PHONE> <ADDR> <#UPDATES>")
	fmt.Println("")
	fmt.Println("APIID 	  means Telegram API id. If you do not have it yet, make it at ")
	fmt.Println("			  https://my.telegram.org/apps")
	fmt.Println("APIHASH   means hashcode of <APIID>. It is published together with API id.")
	fmt.Println("PHONE 	  means already authenticated phonenumber. If you have never")
	fmt.Println("  		  signed in to Telegram with simple shell, authenticate your")
	fmt.Println(" 		  phone through simpleshell first.")
	fmt.Println("ADDR 	  means Telegram server address, e.g., 91.108.56.165:443")
	fmt.Println("#UPDATES  means number of updates to subscribe from the server. Setting")
	fmt.Println(" 		  it as 0 subscribes updates forever.")
}
func main() {
	if len(os.Args) != 6 {
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
	parsed, err := strconv.ParseInt(os.Args[5], 10, 0)
	updateNum := int(parsed)
	if err != nil {
		usage()
		return
	}

	benchtime := time.Now()
	y, mon, d := benchtime.Date()
	h, min, s := benchtime.Clock()
	dateString := fmt.Sprintf("%02d%02d%02d-%02d%02d%02d", (y%100), mon, d, h, min, s)
	lfilename := fmt.Sprintf("%s_%s_benchmark.log", dateString, target)
	bfilename := fmt.Sprintf("%s_%s_benchmark.bch", dateString, target)

	// Set logfile
	lfile, err := os.OpenFile(lfilename, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("error opening file: ", err)
	}
	defer lfile.Close()
	log.SetOutput(lfile)
	//log.SetOutput(os.Stdout)

	// Set slog.Benchfile
	bfile, err := os.OpenFile(bfilename, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("error opening file: ", err)
	}
	defer bfile.Close()
	//bfile := os.Stdout
	slog.SetBenchOutput(bfile)

	slog.Benchln(main, "==== Settings ====")
	slog.Benchln(main, "target: ", target)
	slog.Benchln(main, "updateNum: ", updateNum)
	slog.Benchln(main, "==================")

	// Bench 1: update latency
	benchmarkUpdateLatency(apiid, apihash, phoneNumber, target, updateNum)
}

func benchmarkUpdateLatency(apiid int32, apihash string, phoneNumber string, preferredAddr string, updateNum int) error {
	slog.EnableLogging()
	mm, mconn, err := initBenchmark(apiid, apihash, phoneNumber, preferredAddr)
	if err !=  nil {
		return err
	}

	//slog.DisableLogging()
	slog.Benchln(benchmarkUpdateLatency, "start")
	b := NewUpdateBenchmarker(updateNum)
	mconn.AddUpdateCallback(b)
	b.join.Wait()
	slog.Benchln(benchmarkUpdateLatency, "done")
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

	filtered := filter(floats, func(x float64) bool {return x>500})
	fmean, _ := stats.Mean(floats)
	fmedian, _ := stats.Median(floats)
	fmin, _ := stats.Min(floats)
	fmax, _ := stats.Max(floats)
	fvariance, _ := stats.Variance(floats)
	fstd, _ := stats.StandardDeviation(floats)

	//CAVEAT: it measures latencies by comparing timestamps (unit: sec) embedded in updates and when update callback triggers by NTP clock (unit: ms).)
	//You need to keep in mind two things.)
	// 1. By the unit difference, the timestamped time \\in [0,1) + timestamp)
	//   So remove that noise before analytics. )
	//	   e.g., pick latencies over 0.5, which is an expectation when noise uniformly distributes.)
	// 2. By the time source difference, if Telegram server is not synced with NTP pool, latencies can incorrect.)
	//   So measure latencies from a baseline experiment environment first, and compare it with other environments.)
	slog.Benchln(benchmarkUpdateLatency, "====== RESULTS ======")
	slog.Benchf(benchmarkUpdateLatency, "raw latencies: %s\n", slog.Stringify(b.latencies))
	slog.Benchf(benchmarkUpdateLatency, "raw mean:%f, median:%f, min:%f, max:%f, var:%f, std:%f\n", rmean, rmedian, rmin, rmax, rvariance, rstd)
	slog.Benchf(benchmarkUpdateLatency, "filtered latencies: %s\n", slog.Stringify(filtered))
	slog.Benchf(benchmarkUpdateLatency, "filtered mean:%f, median:%f, min:%f, max:%f, var:%f, std:%f\n", fmean, fmedian, fmin, fmax, fvariance, fstd)
	slog.Benchln(benchmarkUpdateLatency, "=====================")

	// Finish the benchmark
	mm.Finish()
	return nil
}

type condition func(float64) bool

func filter(ints []float64, cond condition) []float64 {
	size := 0
	for _, x := range ints {
		if cond(x) {
			size ++
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
//
//func Float64s(ints []int64) []float64 {
//	floats := make([]float64, len(ints))
//	for i, x := range ints {
//		floats[i] = float64(x)
//	}
//	return floats
//}

func NewUpdateBenchmarker(updateNum int) *updateBenchmarker {
	b := new(updateBenchmarker)
	b.updateCount = 0
	b.updateNum = updateNum
	b.join = sync.WaitGroup{}
	b.join.Add(1)
	return b
}

type updateBenchmarker struct {
	latencies 	[]int64
	updateNum	int
	updateCount	int
	join 		sync.WaitGroup
}

func (b *updateBenchmarker) OnUpdate(u mtproto.MUpdate) {
	tic := ntped.UnixMilli()
	updateDate := int64(u.UpdateDate()) * 1e3	// update date in milli
	diff := tic - updateDate
	if updateDate != 0 {
		slog.Benchf(b, "[tdiff = %d ms] got %s\n", diff, slog.Stringify(u))
		b.latencies = append(b.latencies, diff)
		b.updateCount ++
		if b.updateNum == 0 {
			// never stop
		} else if b.updateCount >= b.updateNum {
			b.join.Done()
		}
	} else {
		slog.Benchf(b, "got %s\n", slog.Stringify(u))
	}
}

func initBenchmark(apiid int32, apihash string, phoneNumber string, preferredAddr string) (*mtproto.MManager, *mtproto.MConn, error){
	const (
		maxNtpRetry = 0
		timeoutInMs = 10000

		appVersion = "0.0.1"
		deviceModel = ""
		systemVersion = ""
		language = ""
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
		configuration, err := mtproto.NewConfiguration(apiid, apihash, appVersion, deviceModel, systemVersion, language, sessionFileHome)
		handleError(err)
		manager, err := mtproto.NewManager(configuration)
		handleError(err)

		if !manager.IsAuthenticated(phoneNumber) {
			out <- fmt.Errorf("auth does NOT exist")
			return
		}
		mconn, err := manager.LoadAuthentication(phoneNumber, preferredAddr)
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

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

