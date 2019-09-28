package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemSaver   chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParserResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	//itemCount := 0
	for {
		result := <-out

		// 打印从out的channel中收到的结果
		for _, item := range result.Items {
			//itemCount++
			//log.Printf("Got item #%d %v", itemCount, item)
			go func(item interface{}) { e.ItemSaver <- item }(item)
		}

		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}
	}
}

func createWorker(in chan Request, out chan ParserResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			// 告诉调度器，这个worker已经可以工作了
			request := <-in
			parserResult, err := worker(request)
			if err != nil {
				continue
			}
			out <- parserResult
		}
	}()
}
