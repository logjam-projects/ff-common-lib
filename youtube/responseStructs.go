package youtube

import "time"

type Response struct {
	ChannelName string                 `json:"channelName"`
	ChannelID   string                 `json:"id"`
	Videos      []InternalStructVideos `json: "videos"`
}

type InternalStructVideos struct {
	Video struct {
		PublishedAt time.Time `json:"publishedAt"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
	} `json:"video"`
	Statistics struct {
		ViewCount     string `json:"viewCount"`
		LikeCount     string `json:"likeCount"`
		DislikeCount  string `json:"dislikeCount"`
		FavoriteCount string `json:"favoriteCount"`
		CommentCount  string `json:"commentCount"`
	} `json:"statistics"`
	Comments []InternalStructComments `json: "comments"`
}

type InternalStructComments struct {
	TopLevelComment struct {
		AuthorDisplayName string `json:"authorDisplayName"`
		AuthorChannelURL  string `json:"authorChannelUrl"`
		AuthorChannelID   struct {
			Value string `json:"value"`
		} `json:"authorChannelId"`
		VideoID      string    `json:"videoId"`
		TextDisplay  string    `json:"textDisplay"`
		TextOriginal string    `json:"textOriginal"`
		CanRate      bool      `json:"canRate"`
		ViewerRating string    `json:"viewerRating"`
		LikeCount    int       `json:"likeCount"`
		PublishedAt  time.Time `json:"publishedAt"`
		UpdatedAt    time.Time `json:"updatedAt"`
	} `json:"topLevelComment"`
	CanReply        bool `json:"canReply"`
	TotalReplyCount int  `json:"totalReplyCount"`
	IsPublic        bool `json:"isPublic"`
}
