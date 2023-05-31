package pull_request

import "time"

type Pull struct {
	Action      string      `json:"action"`
	Number      int         `json:"number"`
	PullRequest PullRequest `json:"pull_request"`
	Repository  Repository  `json:"repository"`
	Sender      Sender      `json:"sender"`
}
type User struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}
type Owner struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}
type Repo struct {
	ID                        int       `json:"id"`
	NodeID                    string    `json:"node_id"`
	Name                      string    `json:"name"`
	FullName                  string    `json:"full_name"`
	Private bool   `json:"private"`
	Owner   Owner  `json:"owner"`
	HTMLURL string `json:"html_url"`
	Description               string    `json:"description"`
	Fork                      bool      `json:"fork"`
	URL                       string    `json:"url"`
	ForksURL                  string    `json:"forks_url"`
	KeysURL                   string    `json:"keys_url"`
	CollaboratorsURL          string    `json:"collaborators_url"`
	TeamsURL                  string    `json:"teams_url"`
	HooksURL                  string    `json:"hooks_url"`
	IssueEventsURL            string    `json:"issue_events_url"`
	EventsURL                 string    `json:"events_url"`
	AssigneesURL              string    `json:"assignees_url"`
	BranchesURL               string    `json:"branches_url"`
	TagsURL                   string    `json:"tags_url"`
	BlobsURL                  string    `json:"blobs_url"`
	GitTagsURL                string    `json:"git_tags_url"`
	GitRefsURL                string    `json:"git_refs_url"`
	TreesURL                  string    `json:"trees_url"`
	StatusesURL               string    `json:"statuses_url"`
	LanguagesURL              string    `json:"languages_url"`
	StargazersURL             string    `json:"stargazers_url"`
	ContributorsURL           string    `json:"contributors_url"`
	SubscribersURL            string    `json:"subscribers_url"`
	SubscriptionURL           string    `json:"subscription_url"`
	CommitsURL                string    `json:"commits_url"`
	GitCommitsURL             string    `json:"git_commits_url"`
	CommentsURL               string    `json:"comments_url"`
	IssueCommentURL           string    `json:"issue_comment_url"`
	ContentsURL               string    `json:"contents_url"`
	CompareURL                string    `json:"compare_url"`
	MergesURL                 string    `json:"merges_url"`
	ArchiveURL                string    `json:"archive_url"`
	DownloadsURL              string    `json:"downloads_url"`
	IssuesURL                 string    `json:"issues_url"`
	PullsURL                  string    `json:"pulls_url"`
	MilestonesURL             string    `json:"milestones_url"`
	NotificationsURL          string    `json:"notifications_url"`
	LabelsURL                 string    `json:"labels_url"`
	ReleasesURL               string    `json:"releases_url"`
	DeploymentsURL            string    `json:"deployments_url"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt                 time.Time `json:"updated_at"`
	PushedAt                  time.Time `json:"pushed_at"`
	GitURL                    string    `json:"git_url"`
	SSHURL                    string    `json:"ssh_url"`
	CloneURL                  string    `json:"clone_url"`
	SvnURL                    string    `json:"svn_url"`
	Homepage                  any       `json:"homepage"`
	Size                      int       `json:"size"`
	StargazersCount           int       `json:"stargazers_count"`
	WatchersCount             int       `json:"watchers_count"`
	Language                  string    `json:"language"`
	HasIssues                 bool      `json:"has_issues"`
	HasProjects               bool      `json:"has_projects"`
	HasDownloads              bool      `json:"has_downloads"`
	HasWiki                   bool      `json:"has_wiki"`
	HasPages                  bool      `json:"has_pages"`
	HasDiscussions            bool      `json:"has_discussions"`
	ForksCount                int       `json:"forks_count"`
	MirrorURL                 any       `json:"mirror_url"`
	Archived                  bool      `json:"archived"`
	Disabled                  bool      `json:"disabled"`
	OpenIssuesCount           int       `json:"open_issues_count"`
	License                   any       `json:"license"`
	AllowForking              bool      `json:"allow_forking"`
	IsTemplate                bool      `json:"is_template"`
	WebCommitSignoffRequired  bool      `json:"web_commit_signoff_required"`
	Topics                    []any     `json:"topics"`
	Visibility                string    `json:"visibility"`
	Forks                     int       `json:"forks"`
	OpenIssues                int       `json:"open_issues"`
	Watchers                  int       `json:"watchers"`
	DefaultBranch             string    `json:"default_branch"`
	AllowSquashMerge          bool      `json:"allow_squash_merge"`
	AllowMergeCommit          bool      `json:"allow_merge_commit"`
	AllowRebaseMerge          bool      `json:"allow_rebase_merge"`
	AllowAutoMerge            bool      `json:"allow_auto_merge"`
	DeleteBranchOnMerge       bool      `json:"delete_branch_on_merge"`
	AllowUpdateBranch         bool      `json:"allow_update_branch"`
	UseSquashPrTitleAsDefault bool      `json:"use_squash_pr_title_as_default"`
	SquashMergeCommitMessage  string    `json:"squash_merge_commit_message"`
	SquashMergeCommitTitle    string    `json:"squash_merge_commit_title"`
	MergeCommitMessage        string    `json:"merge_commit_message"`
	MergeCommitTitle          string    `json:"merge_commit_title"`
}
type Head struct {
	Label string `json:"label"`
	Ref   string `json:"ref"`
	Sha  string `json:"sha"`
	User User   `json:"user"`
	Repo Repo   `json:"repo"`
}
type Base struct {
	Label string `json:"label"`
	Ref   string `json:"ref"`
	Sha  string `json:"sha"`
	User User   `json:"user"`
	Repo Repo   `json:"repo"`
}
type Self struct {
	Href string `json:"href"`
}
type HTML struct {
	Href string `json:"href"`
}
type Issue struct {
	Href string `json:"href"`
}
type Comments struct {
	Href string `json:"href"`
}
type ReviewComments struct {
	Href string `json:"href"`
}
type ReviewComment struct {
	Href string `json:"href"`
}
type Commits struct {
	Href string `json:"href"`
}
type Statuses struct {
	Href string `json:"href"`
}
type Links struct {
	Self           Self           `json:"self"`
	HTML           HTML           `json:"html"`
	Issue          Issue          `json:"issue"`
	Comments       Comments       `json:"comments"`
	ReviewComments ReviewComments `json:"review_comments"`
	ReviewComment  ReviewComment  `json:"review_comment"`
	Commits        Commits        `json:"commits"`
	Statuses       Statuses       `json:"statuses"`
}
type PullRequest struct {
	URL                 string    `json:"url"`
	ID                  int       `json:"id"`
	NodeID              string    `json:"node_id"`
	HTMLURL             string    `json:"html_url"`
	DiffURL             string    `json:"diff_url"`
	PatchURL            string    `json:"patch_url"`
	IssueURL            string    `json:"issue_url"`
	Number              int       `json:"number"`
	State               string    `json:"state"`
	Locked              bool      `json:"locked"`
	Title             string `json:"title"`
	User              User   `json:"user"`
	Body              any    `json:"body"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	ClosedAt            any       `json:"closed_at"`
	MergedAt            any       `json:"merged_at"`
	MergeCommitSha      any       `json:"merge_commit_sha"`
	Assignee            any       `json:"assignee"`
	Assignees           []any     `json:"assignees"`
	RequestedReviewers  []any     `json:"requested_reviewers"`
	RequestedTeams      []any     `json:"requested_teams"`
	Labels              []any     `json:"labels"`
	Milestone           any       `json:"milestone"`
	Draft               bool      `json:"draft"`
	CommitsURL          string    `json:"commits_url"`
	ReviewCommentsURL   string    `json:"review_comments_url"`
	ReviewCommentURL    string    `json:"review_comment_url"`
	CommentsURL         string    `json:"comments_url"`
	StatusesURL       string `json:"statuses_url"`
	Head              Head   `json:"head"`
	Base              Base   `json:"base"`
	Links             Links  `json:"_links"`
	AuthorAssociation string `json:"author_association"`
	AutoMerge           any       `json:"auto_merge"`
	ActiveLockReason    any       `json:"active_lock_reason"`
	Merged              bool      `json:"merged"`
	Mergeable           any       `json:"mergeable"`
	Rebaseable          any       `json:"rebaseable"`
	MergeableState      string    `json:"mergeable_state"`
	MergedBy            any       `json:"merged_by"`
	Comments            int       `json:"comments"`
	ReviewComments      int       `json:"review_comments"`
	MaintainerCanModify bool      `json:"maintainer_can_modify"`
	Commits             int       `json:"commits"`
	Additions           int       `json:"additions"`
	Deletions           int       `json:"deletions"`
	ChangedFiles        int       `json:"changed_files"`
}
type Repository struct {
	ID                       int       `json:"id"`
	NodeID                   string    `json:"node_id"`
	Name                     string    `json:"name"`
	FullName                 string    `json:"full_name"`
	Private bool   `json:"private"`
	Owner   Owner  `json:"owner"`
	HTMLURL string `json:"html_url"`
	Description              string    `json:"description"`
	Fork                     bool      `json:"fork"`
	URL                      string    `json:"url"`
	ForksURL                 string    `json:"forks_url"`
	KeysURL                  string    `json:"keys_url"`
	CollaboratorsURL         string    `json:"collaborators_url"`
	TeamsURL                 string    `json:"teams_url"`
	HooksURL                 string    `json:"hooks_url"`
	IssueEventsURL           string    `json:"issue_events_url"`
	EventsURL                string    `json:"events_url"`
	AssigneesURL             string    `json:"assignees_url"`
	BranchesURL              string    `json:"branches_url"`
	TagsURL                  string    `json:"tags_url"`
	BlobsURL                 string    `json:"blobs_url"`
	GitTagsURL               string    `json:"git_tags_url"`
	GitRefsURL               string    `json:"git_refs_url"`
	TreesURL                 string    `json:"trees_url"`
	StatusesURL              string    `json:"statuses_url"`
	LanguagesURL             string    `json:"languages_url"`
	StargazersURL            string    `json:"stargazers_url"`
	ContributorsURL          string    `json:"contributors_url"`
	SubscribersURL           string    `json:"subscribers_url"`
	SubscriptionURL          string    `json:"subscription_url"`
	CommitsURL               string    `json:"commits_url"`
	GitCommitsURL            string    `json:"git_commits_url"`
	CommentsURL              string    `json:"comments_url"`
	IssueCommentURL          string    `json:"issue_comment_url"`
	ContentsURL              string    `json:"contents_url"`
	CompareURL               string    `json:"compare_url"`
	MergesURL                string    `json:"merges_url"`
	ArchiveURL               string    `json:"archive_url"`
	DownloadsURL             string    `json:"downloads_url"`
	IssuesURL                string    `json:"issues_url"`
	PullsURL                 string    `json:"pulls_url"`
	MilestonesURL            string    `json:"milestones_url"`
	NotificationsURL         string    `json:"notifications_url"`
	LabelsURL                string    `json:"labels_url"`
	ReleasesURL              string    `json:"releases_url"`
	DeploymentsURL           string    `json:"deployments_url"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
	PushedAt                 time.Time `json:"pushed_at"`
	GitURL                   string    `json:"git_url"`
	SSHURL                   string    `json:"ssh_url"`
	CloneURL                 string    `json:"clone_url"`
	SvnURL                   string    `json:"svn_url"`
	Homepage                 any       `json:"homepage"`
	Size                     int       `json:"size"`
	StargazersCount          int       `json:"stargazers_count"`
	WatchersCount            int       `json:"watchers_count"`
	Language                 string    `json:"language"`
	HasIssues                bool      `json:"has_issues"`
	HasProjects              bool      `json:"has_projects"`
	HasDownloads             bool      `json:"has_downloads"`
	HasWiki                  bool      `json:"has_wiki"`
	HasPages                 bool      `json:"has_pages"`
	HasDiscussions           bool      `json:"has_discussions"`
	ForksCount               int       `json:"forks_count"`
	MirrorURL                any       `json:"mirror_url"`
	Archived                 bool      `json:"archived"`
	Disabled                 bool      `json:"disabled"`
	OpenIssuesCount          int       `json:"open_issues_count"`
	License                  any       `json:"license"`
	AllowForking             bool      `json:"allow_forking"`
	IsTemplate               bool      `json:"is_template"`
	WebCommitSignoffRequired bool      `json:"web_commit_signoff_required"`
	Topics                   []any     `json:"topics"`
	Visibility               string    `json:"visibility"`
	Forks                    int       `json:"forks"`
	OpenIssues               int       `json:"open_issues"`
	Watchers                 int       `json:"watchers"`
	DefaultBranch            string    `json:"default_branch"`
}
type Sender struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}
