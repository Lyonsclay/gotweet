
// TWEET LOOKUP
// GET/POST /2/tweets
// DELETE/GET /2/tweets/{id}
// GET /2/tweets/{id}/retweeted_by
// GET /2/tweets/counts/recent
// GET /2/tweets/counts/all
// SEARCH TWEETS
// GET /2/tweets/search/recent
// GET /2/tweets/search/all
// TIMELINES
// GET /2/users/{id}/tweets
// GET /2/users/{id}/mentions
// GET /2/lists/{id}/tweets
// FILTERED STREAM
// GET /2/tweets/search/stream
// GET/POST /2/tweets/search/stream/rules
// SAMPLED STREAM
// GET /2/tweets/sample/stream
// LIKES
// POST /2/users/{id}/likes
// DELETE /2/users/{id}/likes/{tweet_id}
// GET /2/tweets/{id}/liking_users
// HIDE REPLIES
// PUT /2/tweets/{id}/hidden
// USER LOOKUP
// GET /2/users
// GET /2/users/{id}
// GET /2/users/by
// GET /2/users/me
// GET /2/users/by/username/{username}
// DELETE /2/users/{source_user_id}/muting/{target_user_id}
// GET/POST /2/users/{id}/muting
// GET /2/users/{id}/list_memberships
// GET /2/users/{id}/owned_lists
// GET/POST /2/users/{id}/pinned_lists
// DELETE /2/users/{id}/pinned_lists/{list_id}
// GET /2/users/{id}/liked_tweets
// POST /2/users/{id}/retweets
// DELETE /2/users/{id}/retweets/{source_tweet_id}
// FOLLOWS
// GET /2/users/{id}/followers
// GET/POST /2/users/{id}/following
// DELETE /2/users/{source_user_id}/following/{target_user_id}
// GET/POST /2/users/{id}/followed_lists
// DELETE /2/users/{id}/followed_lists/{list_id}
// GET /2/lists/{id}/followers
// BLOCKS
// GET/POST /2/users/{id}/blocking
// DELETE /2/users/{source_user_id}/blocking/{target_user_id}
// SPACES
// GET /2/spaces/{id}
// GET /2/spaces
// GET /2/spaces/by/creator_ids
// GET /2/spaces/search
// GET /2/spaces/{id}/tweets
// GET /2/spaces/{id}/buyers
// COMPLIANCE
// GET/POST /2/compliance/jobs
// GET /2/compliance/jobs/{id}
// LISTS
// POST /2/lists
// DELETE/GET/PUT /2/lists/{id}
// GET/POST /2/lists/{id}/members
// DELETE /2/lists/{id}/members/{user_id}

// _ Tweets
//
// __ Tweets Lookup
// GET /2/tweets
// - Retrieve multiple Tweets with a list of IDs
// GET /2/tweets/:id
// - Retrieve a single Tweet with an ID
// -- Manage tweets
// POST /2/tweets
// - Post a Tweet
// DELETE /2/tweets/:id
// - Delete a Tweet
//
// __ Timelines lookup
// ___ User Tweet timeline
// GET /2/users/:id/tweets
// - Returns most recent Tweets composed a specified user ID
// ___ User mention timeline
// GET /2/users/:id/mentions
// - Returns most recent Tweets mentioning a specified user ID
//
// __ Search Tweets
// ___ Recent Search
// GET /2/tweets/search/recent
// - Search for Tweets published in the last 7 days
// ___ Full-archive search
// GET /2/tweets/search/all
// - Search the full archive of Tweets
// - Only available to those with Academic Research access
//
// __ Tweet counts
// ___ Recent Tweet counts
// GET /2/tweets/counts/recent
// - Receive a count of Tweets that match a query in the last 7 days
// ___ Full-archive Tweet counts
// GET /2/tweets/counts/all
// - Only available to those with Academic Research access
// - Receive a count of Tweets that match a query
// GET /2/tweets/counts/all
// - Search the full archive of Tweets
//
// __ Filtered stream
// POST /2/tweets/search/stream/rules
// - Add or delete rules from your stream
// GET /2/tweets/search/stream/rules
// - Retrieve your stream's rules
// GET /2/tweets/search/stream
// - Connect to the stream
// GET /2/tweets/sample/stream
// - Streams about 1% of all Tweets in real-time.
// - OAuth 2.0 App-only
//
//
// _ Retweets
//
// __ Retweets lookup
// GET /2/users/:id/retweeted_by
// - Users who have Retweeted a Tweet
//
// __ Manage Retweets
// POST /2/users/:id/retweets
// - Allows a user ID to Retweet a Tweet
// DELETE /2/users/:id/retweets/:source_tweet_id
// - Allows a user ID to undo a Retweet
//
//
// _ Likes
//
// __ Likes lookup
// GET /2/tweets/:id/liking_users
// - Users who have liked a Tweet
// GET /2/users/:id/liked_tweets
// - Tweets liked by a user
//
// __Manage Likes
// POST /2/users/:id/likes
// - Allows a user ID to like a Tweet
// DELETE /2/users/:id/likes/:tweet_id
// - Allows a user ID to unlike a Tweet
//
// __ Hide Replies
// PUT /2/tweets/:id/hidden
// - Hides or unhides a reply to a Tweet.
//
//
// Please note: The COVID-19 stream application form is currently closed.
// We currently do not have plans to reopen the application process.
// __ COVID-19 stream
// GET /labs/1/tweets/stream/covid19
// - OAuth 2.0 App-Only
// GET /labs/1/tweets/stream/compliance
// - OAuth 2.0 App-Only
//
// _ Users
//
// __ Users lookup
// GET /2/users
// - Retrieve multiple users with usernames
// GET /2/users/:id
// - Retrieve a single user with an ID
// GET /2/users/by
// - Retrieve multiple users with usernames
// GET /2/users/by/username/:username
// - Retrieve a single user with a username
// GET /2/users/me
// - Returns the information about an authorized user
//
// __ Follows lookup
// GET /2/users/:id/following
// GET /2/users/:id/followers
//
// __ Manage follows
// POST /2/users/:id/following
// - Allows a user ID to follow another user
// DELETE /2/users/:source_user_id/following/:target_user_id
// - Allows a user ID to unfollow another user
//
// __ Blocks lookup
// GET /2/users/:id/blocking
// - Returns a list of users who are blocked by the specified user ID
//
// __Manage blocks
// POST /2/users/:id/blocking
// - Allows a user ID to block another user
// DELETE /2/users/:source_user_id/blocking/:target_user_id
// - Allows a user ID to unblock another user
//
// __ Mutes lookup
// GET /2/users/:id/muting
// - Returns a list of users who are muted by the specified user ID
//
// __ Manage mutes
// POST /2/users/:id/muting
// - Allows a user ID to mute another user
// DELETE /2/users/:source_user_id/muting/:target_user_id
// - Allows a user ID to unmute another user
//
