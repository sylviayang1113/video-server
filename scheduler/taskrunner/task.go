package taskrunner

import (
	"errors"
	"log"
	"github.com/interesting1113/video-server/video_server/scheduler/dbops"
)


func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("Video clear dispatcher error: %v", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("All tasks finished")
	}

	for _, id := range res {
		dc <- id
	}

	return nil
}


func VideoClearExecutor(dc dataChan) error {
	

}