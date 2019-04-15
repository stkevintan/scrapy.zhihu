package api

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"

	"github.com/ddliu/go-httpclient"
)

type TopicDescr struct {
	ID      string
	Limit   int
	AfterID string
	Include []string
}

//NewTopicDescr get the topicDescr instance
func NewTopicDescr(id string, limit int, afterID string) *TopicDescr {
	topic := &TopicDescr{id, limit, afterID, nil}
	//事实上，不需要恶心信息
	//reference: https://www.twblogs.net/a/5c6e516abd9eee5c86dcee6e
	topic.Include = []string{
		// "data[?(target.type=topic_sticky_module)].target.data[?(target.type=answer)].target.content",
		// "relationship.is_authorized",
		// "is_author",
		// "voting",
		// "is_thanked",
		// "is_nothelp;data[?(target.type=topic_sticky_module)].target.data[?(target.type=answer)].target.is_normal",
		// "comment_count",
		// "voteup_count",
		// "content",
		// "relevant_info",
		// "excerpt.author.badge[?(type=best_answerer)].topics;data[?(target.type=topic_sticky_module)].target.data[?(target.type=article)].target.content",
		// "voteup_count",
		// "comment_count",
		// "voting",
		// "author.badge[?(type=best_answerer)].topics;data[?(target.type=topic_sticky_module)].target.data[?(target.type=people)].target.answer_count",
		// "articles_count",
		// "gender",
		// "follower_count",
		// "is_followed",
		// "is_following",
		// "badge[?(type=best_answerer)].topics;data[?(target.type=answer)].target.annotation_detail",
		// "content",
		// "hermes_label",
		// "is_labeled",
		// "relationship.is_authorized",
		// "is_author",
		// "voting",
		// "is_thanked",
		// "is_nothelp;data[?(target.type=answer)].target.author.badge[?(type=best_answerer)].topics;data[?(target.type=article)].target.annotation_detail",
		// "content",
		// "hermes_label",
		// "is_labeledj",
		// "author.badge[?(type=best_answerer)].topics;data[?(target.type=question)].target.annotation_detail",
		// "comment_count;",
	}
	return topic
}

//TopicList get the topic list
func (topic *TopicDescr) TopicList() ([]byte, error) {

	urlStr := fmt.Sprintf("https://www.zhihu.com/api/v4/topics/%s/feeds/top_activity", topic.ID)

	limit := topic.Limit
	if limit <= 0 {
		limit = 20
	}

	res, err := httpclient.Get(urlStr, url.Values{
		"limit":    {strconv.Itoa(limit)},
		"include":  {strings.Join(topic.Include, ",")},
		"after_id": {topic.AfterID},
	})

	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	defer func() {
		_ = res.Body.Close()
	}()
	return bodyBytes, nil
}
