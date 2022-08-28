package bot

import (
	"discordbot/config"
	"github.com/dghubble/go-twitter/twitter"
	"strconv"
)

func GetTimeline(username string) (tweet twitter.Tweet, err error) {
	client := twitter.NewClient(config.CreateTwitterClient())
	id, err := strconv.ParseInt(username, 10, 64)
	if err == nil {
		timeline, httpresponse, err2 := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
			UserID: id,
		})
		if tests(err2, httpresponse.StatusCode, len(timeline)) {
			return
		}
		tweet = timeline[0]
		return
	}

	timeline, httpresponse, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: username,
	})
	if tests(err, httpresponse.StatusCode, len(timeline)) {
		return
	}
	tweet = timeline[0]
	return
}

func tests(err error, responsecode int, length int) bool {
	if err != nil || responsecode != 200 || length < 1 {
		return true
	}
	return false
}
