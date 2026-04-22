package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

// func doneAwareStream(in In, done In) Out {
// 	out := make(chan interface{})
// 	go func() {
// 		defer close(out)
// 		for {
// 			select {
// 			case <-done:
// 				return
// 			case v, ok := <-in:
// 				if !ok {
// 					return
// 				}
// 				select {
// 				case out <- v:
// 				case <-done:
// 					return
// 				}
// 			}
// 		}
// 	}()
// 	return out
// }

// func safeWrapper(done In, in In, out Bi) {
// 	defer close(out)

// 	for {
// 		select {
// 		case <-done:
// 			for range in {
// 			}
// 			return
// 		case v, ok := <-in:
// 			if !ok {
// 				return
// 			}
// 			select {
// 			case out <- v:
// 			case <-done:
// 				for range in {
// 				}
// 				return
// 			}
// 		}
// 	}
// }

// func ExecutePipeline(in In, done In, stages ...Stage) Out {
// 	out := in

// 	for _, stage := range stages {
// 		inCh := out
// 		stageOut := stage(doneAwareStream(inCh, done))
// 		// stageOut := stage(inCh)
// 		// out = stageOut

// 		next := make(Bi)

// 		go safeWrapper(done, stageOut, next)

// 		out = next
// 	}

// 	return out
// }

func withDone(in In, done In) In {
	out := make(chan interface{})

	go func() {
		defer close(out)

		for {
			select {
			case <-done:
				for range in {
				}
				return

			case v, ok := <-in:
				if !ok {
					return
				}

				select {
				case out <- v:
				case <-done:
					for range in {
					}
					return
				}
			}
		}
	}()

	return out
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in

	for _, stage := range stages {
		out = stage(withDone(out, done))
	}

	return out
}
