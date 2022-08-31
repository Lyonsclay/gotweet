package gotweet

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

func getConn() *pgx.Conn {
	cs, exists := os.LookupEnv("PG_TWITTER_LOCAL")
	if !exists {
		log.Fatal("You need to provide a connection string in an environment variable\n \"PG_TWITTER_LOCAL\"")
	}
	conn, err := pgx.Connect(context.Background(), cs)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func getOid(name string) uint32 {
	var oid uint32
	conn := getConn()
	err := conn.QueryRow(context.Background(), "select oid from pg_type where typname=$1;", name).Scan(&oid)
	if err != nil {
		log.Fatalf("Could not find  pgtype %s - %v\n", name, err)
	}
	return oid
}

func getArrayOid(name string) uint32 {
	var oid uint32
	conn := getConn()
	err := conn.QueryRow(context.Background(), "select oid from pg_type where typname=$1;", name).Scan(&oid)

	if err != nil {
		log.Fatal(err)
	}
	return oid
}

func getSafeOid(ci *pgtype.ConnInfo, name string) uint32 {
	h, tf := ci.DataTypeForName(name)
	if tf {
		return h.OID
	}
	return getOid(name)
}

func getSafeArrayOid(ci *pgtype.ConnInfo, name string) uint32 {
	h, tf := ci.DataTypeForName(name)
	if tf {
		return h.OID
	}
	return getArrayOid(name)
}

func RegisterHashtag(ci *pgtype.ConnInfo) (*pgtype.CompositeType, uint32) {
	name := "user_hashtag"
	oid := getSafeOid(ci, name)
	hashtag, err := pgtype.NewCompositeType(name, []pgtype.CompositeTypeField{
		{Name: "start", OID: pgtype.Int2OID},
		{Name: "end", OID: pgtype.Int2OID},
		{Name: name, OID: pgtype.VarcharOID},
	}, ci)
	if err != nil {
		fmt.Println(err)
	}
	ci.RegisterDataType(pgtype.DataType{Value: hashtag, Name: hashtag.TypeName(), OID: oid})
	return hashtag, oid
}

func RegisterHashtags(ci *pgtype.ConnInfo) (*pgtype.ArrayType, uint32) {
	name := "_user_hashtag"
	oid := getSafeOid(ci, name)
	h, hOid := RegisterHashtag(ci)
	hashtags := pgtype.NewArrayType(name, hOid, func() pgtype.ValueTranscoder { return h })
	ci.RegisterDataType(pgtype.DataType{Value: hashtags, Name: hashtags.TypeName(), OID: oid})
	return hashtags, oid
}


func RegisterMention(ci *pgtype.ConnInfo) (*pgtype.CompositeType, uint32) {
	name := "mention"
	oid := getSafeOid(ci, name)
	mention, err := pgtype.NewCompositeType(name, []pgtype.CompositeTypeField{
		{Name: "start", OID: pgtype.Int2OID},
		{Name: "end", OID: pgtype.Int2OID},
		{Name: "usernae", OID: pgtype.VarcharOID},
	}, ci)
	if err != nil {
		log.Fatalf("RegisterMention create pgtype.NewCompositeType failed - %v\n", err)
	}
	ci.RegisterDataType(pgtype.DataType{Value: mention, Name: mention.TypeName(), OID: oid})
	return mention, oid
}


func RegisterMentions(ci *pgtype.ConnInfo) (*pgtype.ArrayType, uint32) {
	name := "_mention"
	oid := getSafeOid(ci, name)
	m, mOid := RegisterMention(ci)
	mentions := pgtype.NewArrayType(name, mOid, func() pgtype.ValueTranscoder { return m })
	ci.RegisterDataType(pgtype.DataType{Value: mentions, Name: mentions.TypeName(), OID: oid})
	return mentions, oid
}


func RegisterCashtag(ci *pgtype.ConnInfo) (*pgtype.CompositeType, uint32) {
	name := "user_cashtag"
	oid := getSafeOid(ci, name)
	cashtag, err := pgtype.NewCompositeType(name, []pgtype.CompositeTypeField{
		{Name: "start", OID: pgtype.Int2OID},
		{Name: "end", OID: pgtype.Int2OID},
		{Name: "cashtag", OID: pgtype.VarcharOID},
	}, ci)
	if err != nil {
		log.Fatal(err)
	}
	ci.RegisterDataType(pgtype.DataType{Value: cashtag, Name: cashtag.TypeName(), OID: oid})
	return cashtag, oid
}


func RegisterCashtags(ci *pgtype.ConnInfo) (*pgtype.ArrayType, uint32) {
	name := "_user_cashtag"
	oid := getSafeArrayOid(ci, name)
	c, cOid := RegisterCashtag(ci)
	cashtags := pgtype.NewArrayType(name, cOid, func() pgtype.ValueTranscoder { return c })
	ci.RegisterDataType(pgtype.DataType{Value: cashtags, Name: cashtags.TypeName(), OID: oid})
	return cashtags, oid
}

func RegisterUserUrlExpansion(ci *pgtype.ConnInfo) (*pgtype.CompositeType, uint32) {
	name := "user_url_expansion"
	oid := getSafeOid(ci, name)
	urlExpansion, err := pgtype.NewCompositeType(name, []pgtype.CompositeTypeField{
		{Name: "start", OID: pgtype.Int2OID},
		{Name: "end", OID: pgtype.Int2OID},
		{Name: "url", OID: pgtype.VarcharOID},
		{Name: "expanded_url", OID: pgtype.VarcharOID},
		{Name: "display_url", OID: pgtype.VarcharOID},
	}, ci)
	if err != nil {
		fmt.Println(err)
	}
	ci.RegisterDataType(pgtype.DataType{Value: urlExpansion, Name: urlExpansion.TypeName(), OID: oid})
	return urlExpansion, oid
}

func RegisterUserUrlExpansions(ci *pgtype.ConnInfo) (*pgtype.ArrayType, uint32) {
	name := "_user_url_expansion"
	oid := getSafeArrayOid(ci, "_user_url_expansion")
	c, uOid := RegisterUserUrlExpansion(ci)
	urls := pgtype.NewArrayType(name, uOid, func() pgtype.ValueTranscoder { return c })
	ci.RegisterDataType(pgtype.DataType{Value: urls, Name: urls.TypeName(), OID: oid})
	return urls, oid
}


func RegisterUserEntitiesUrl(ci *pgtype.ConnInfo) (*pgtype.CompositeType, uint32) {
	name := "user_entities_url"
	oid := getSafeOid(ci, name)
	_, urlsOid := RegisterUserUrlExpansions(ci)
	entitiesUrl, err := pgtype.NewCompositeType(name, []pgtype.CompositeTypeField{{Name: "urls", OID: urlsOid}}, ci)
	if err != nil {
		log.Fatal(err)
	}
	ci.RegisterDataType(pgtype.DataType{
		Value: entitiesUrl,
		Name:  entitiesUrl.TypeName(),
		OID:   oid,
	})
	return entitiesUrl, oid
}


func RegisterUserDescription(ci *pgtype.ConnInfo) (*pgtype.CompositeType, uint32) {
	name := "user_description"
	oid := getSafeOid(ci, name)
	_, euOid := RegisterUserUrlExpansions(ci)
	_, hOid := RegisterHashtags(ci)
	_, mOid := RegisterMentions(ci)
	_, cOid := RegisterCashtags(ci)
	userDescription, err := pgtype.NewCompositeType("user_description", []pgtype.CompositeTypeField{
		{Name: "urls", OID: euOid},
		{Name: "hashtags", OID: hOid},
		{Name: "mentions", OID: mOid},
		{Name: "cashtags", OID: cOid},
	}, ci)
	if err != nil {
		log.Fatal(err)
	}
	ci.RegisterDataType(pgtype.DataType{Value: userDescription, Name: userDescription.TypeName(), OID: oid})
	return userDescription, oid
}

func RegisterUserEntities(ci *pgtype.ConnInfo) (*pgtype.CompositeType, uint32) {
	name := "user_entities"
	oid := getSafeOid(ci, name)
	_, eusOid := RegisterUserEntitiesUrl(ci)
	_, edOid := RegisterUserDescription(ci)
	entities, err := pgtype.NewCompositeType(name, []pgtype.CompositeTypeField{
		{Name: "url", OID: eusOid},
		{Name: "description", OID: edOid},
	}, ci)
	if err != nil {
		log.Fatalf("Create entities NewCompositeType  failed - %v\n", err)
	}
	ci.RegisterDataType(pgtype.DataType{Value: entities, Name: entities.TypeName(), OID: oid})
	return entities, oid
}


func RegisterUserPublicMetrics(ci *pgtype.ConnInfo) (*pgtype.CompositeType, uint32) {
	name := "user_public_metrics"
	oid := getSafeOid(ci, name)
	pm, err := pgtype.NewCompositeType("user_public_metrics", []pgtype.CompositeTypeField{
		{Name: "followers_count", OID: pgtype.Int8OID},
		{Name: "following_count", OID: pgtype.Int8OID},
		{Name: "tweet_count", OID: pgtype.Int8OID},
		{Name: "listed_count", OID: pgtype.Int8OID},
	}, ci)
	if err != nil {
		log.Fatal(err)
	}
	ci.RegisterDataType(pgtype.DataType{Value: pm, Name: pm.TypeName(), OID: oid})
	return pm, oid
}

func RegisterWithheld(ci *pgtype.ConnInfo) (*pgtype.CompositeType, uint32) {
	wOid := getSafeOid(ci, "user_withheld")
	wsOid := getSafeOid(ci, "withheld_scope")
	scope := pgtype.NewEnumType("withheld_scope", []string{"tweet", "user"})
	ci.RegisterDataType(pgtype.DataType{Value: scope, Name: scope.TypeName(), OID: wsOid})
	wType := []pgtype.CompositeTypeField{
		{Name: "country_codes", OID: pgtype.VarcharArrayOID},
		{Name: "scope", OID: wsOid},
	}
	w, err := pgtype.NewCompositeType("withheld", wType, ci)
	if err != nil {
		log.Fatal(err)
	}
	ci.RegisterDataType(pgtype.DataType{Value: w, Name: w.TypeName(), OID: wOid})
	return w, wOid
}

func (src UserWithheld) Encode(ci *pgtype.ConnInfo) interface{} {
	if src.Scope != nil || len(src.CountryCodes) > 0 {
		return []interface{}{&src.CountryCodes, &src.Scope}
	} else {
		return nil
	}
}

func (src UserEntities) Encode(ci *pgtype.ConnInfo) interface{} {
	var hashtags []interface{}
	for _, h := range src.Description.Hashtags {
		ht := []interface{}{&h.Start, &h.End, &h.Hashtag}
		hashtags = append(hashtags, ht)
	}
	var mentions []interface{}
	for _, m := range src.Description.Mentions {
		mention := []interface{}{&m.Start, &m.End, &m.Username}
		mentions = append(mentions, mention)
	}
	var cashtags []interface{}
	for _, c := range src.Description.Cashtags {
		cashtag := []interface{}{&c.Start, &c.End, &c.Cashtag}
		cashtags = append(cashtags, cashtag)
	}
	var urls []interface{}
	for _, u := range src.Url.Urls {
		url := []interface{}{&u.Start, &u.End, &u.Url, &u.ExpandedUrl, &u.DisplayUrl}
		urls = append(urls, url)
	}
	var descriptionUrls []interface{}
	for _, u := range src.Description.Urls {
		url := []interface{}{&u.Start, &u.End, &u.Url, &u.ExpandedUrl, &u.DisplayUrl}
		descriptionUrls = append(descriptionUrls, url)
	}

	changes := len(urls) + len(hashtags) + len(cashtags) + len(mentions) + len(descriptionUrls)
	uFace := []interface{}{urls}
	dFace := []interface{}{descriptionUrls, hashtags, mentions, cashtags}
	if changes > 0 {
		return []interface{}{uFace, dFace}
	} else {
		return nil
	}
}

func (src UserPublicMetrics) Encode(ci *pgtype.ConnInfo) interface{} {
	fc := src.FollowersCount
	ic := src.FollowingCount
	tc := src.TweetCount
	lc := src.ListedCount
	if fc != nil || ic != nil || tc != nil || lc != nil {
		return []interface{}{&fc, &ic, &tc, &lc}
	} else {
		return nil
	}
}

func prepUserColumns(ci *pgtype.ConnInfo, b *User) ([]interface{}, error) {
	var values [14]interface{}
	values[0] = b.Id
	values[1] = b.Name
	values[2] = b.Username
	values[3] = b.CreatedAt
	values[4] = b.Protected
	values[5] = b.Withheld.Encode(ci)
	values[6] = b.Location
	values[7] = b.Url
	values[8] = b.Description
	values[9] = b.Verified
	values[10] = b.Entities.Encode(ci)
	values[11] = b.PublicImageUrl
	values[12] = b.PublicMetrics.Encode(ci)
	values[13] = b.PinnedTweetId

	return values[:], error(nil)
}

func InsertBlockedUsers(c string, b []User) error{
	conn, err := pgx.Connect(context.Background(), c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	ci := conn.ConnInfo()
	RegisterUserEntities(ci)
	RegisterWithheld(ci)
	RegisterUserPublicMetrics(ci)
	q := INSERT_INTO_USER_BLOCKS
	batch := &pgx.Batch{}
	for _, u := range b {
		cols, err := prepUserColumns(pgtype.NewConnInfo(), &u)
		if err != nil {
			log.Printf("prepUserColumns(u) encountered an error - %s\n", err)
			return err
		}
		batch.Queue(q, cols...)
	}
	br := conn.SendBatch(context.Background(), batch)
	defer br.Close()
	rows, err := br.Query()
	fmt.Printf("here are the rows -- %v\n", rows)
	if err != nil {
		log.Fatalf("InsertUsers error - %s\n", err)
	}

	return err
}

func CopyFromBlocks(c string, b []User) (int64, error) {
	conn, err := pgx.Connect(context.Background(), c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	ci := conn.ConnInfo()
	RegisterUserEntities(ci)
	RegisterWithheld(ci)
	RegisterUserPublicMetrics(ci)
	var inputRows [][]interface{}
	for _, u  := range b {
		cols, err := prepUserColumns(ci, &u)
		if err != nil {
			log.Printf("prepUserColumns(u) encountered an error - %s\n", err)
			return 0, err
		}
		inputRows = append(inputRows, cols)
	}
	columns := []string{"twitter_id", "name", "username", "created_at", "protected", "withheld", "location", "url", "description", "verified", "entities", "public_image_url", "public_metrics", "pinned_tweet_id"}
	copyCount, err := conn.CopyFrom(context.Background(), pgx.Identifier{"users","blocks"}, columns, pgx.CopyFromRows(inputRows))
	if err != nil {
		log.Fatalf("InsertUsers error - %s\n", err)
	}

	return copyCount, err
}
