package youtube

import "time"

type Status struct {
	State string `json:"state,omitempty"`
}

// Youtube core response

type UserID struct {
	Items []struct {
		ID string `json:"id"`
	} `json:"items"`
}

type PlayListByDateResponse struct {
	Items []struct {
		ID struct {
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
		} `json:"snippet"`
	} `json:"items"`
}

type VideoDetailsAndStats struct {
	Items []struct {
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
			ChannelID   string    `json:"channelId"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
		} `json:"snippet"`
		Statistics struct {
			ViewCount     string `json:"viewCount"`
			LikeCount     string `json:"likeCount"`
			DislikeCount  string `json:"dislikeCount"`
			FavoriteCount string `json:"favoriteCount"`
			CommentCount  string `json:"commentCount"`
		} `json:"statistics"`
	} `json:"items"`
}

type VideoComments struct {
	Kind          string `json:"kind"`
	Etag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	PageInfo      struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Kind    string `json:"kind"`
		Etag    string `json:"etag"`
		ID      string `json:"id"`
		Snippet struct {
			VideoID         string `json:"videoId"`
			TopLevelComment struct {
				Kind    string `json:"kind"`
				Etag    string `json:"etag"`
				ID      string `json:"id"`
				Snippet struct {
					AuthorDisplayName     string `json:"authorDisplayName"`
					AuthorProfileImageURL string `json:"authorProfileImageUrl"`
					AuthorChannelURL      string `json:"authorChannelUrl"`
					AuthorChannelID       struct {
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
				} `json:"snippet"`
			} `json:"topLevelComment"`
			CanReply        bool `json:"canReply"`
			TotalReplyCount int  `json:"totalReplyCount"`
			IsPublic        bool `json:"isPublic"`
		} `json:"snippet"`
	} `json:"items"`
}
