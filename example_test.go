package genidbench

import (
	crand "crypto/rand"
	"fmt"
	"time"

	snowflake "github.com/bwmarrin/snowflake"
	sandflake "github.com/celrenheit/sandflake"
	uuiddraft "github.com/coding-socks/uuiddraft"
	gofrs_uuid "github.com/gofrs/uuid"
	google_uuid "github.com/google/uuid"
	shortuuid "github.com/lithammer/shortuuid/v3"
	nanoid "github.com/matoous/go-nanoid/v2"
	ulid "github.com/oklog/ulid/v2"
	xid "github.com/rs/xid"
	ksuid "github.com/segmentio/ksuid"
	sonyflake "github.com/sony/sonyflake"
	hashid "github.com/speps/go-hashids/v2"
	mysql "gorm.io/driver/mysql"
	gorm "gorm.io/gorm"
)

func ExampleUUIDv1() {
	fmt.Println(google_uuid.Must(google_uuid.NewUUID()).String())
	// Output: ddb39472-ddba-11eb-8b44-149d9980a4d2
}
func ExampleUUIDv3() {
	fmt.Println(gofrs_uuid.NewV3(gofrs_uuid.NamespaceDNS, "example.com").String())
	// Output: 9073926b-929f-31c2-abc9-fad77ae3e8eb
}

func ExampleUUIDv4() {
	fmt.Println(google_uuid.New().String())
	// Output: 031cba4b-cbfc-463f-afe8-a50af36f3eb3
}

func ExampleUUIDv5() {
	fmt.Println(gofrs_uuid.NewV5(gofrs_uuid.NamespaceDNS, "example.com").String())
	// Output: cfbff0d1-9375-5685-968c-48ce8b15ae17
}

func ExampleUUIDv6() {
	fmt.Println(uuiddraft.Must(uuiddraft.NewV6()).String())
	// Output: 1ebddbad-db3a-6296-8000-8da19654df48
}

func ExampleUUIDv7() {
	fmt.Println(uuiddraft.Must(uuiddraft.NewV7()).String())
	// Output: 060e348b-5c1b-7943-8000-c414add1c8d1
}

func ExampleULID() {
	fmt.Println(ulid.MustNew(uint64(time.Now().UnixNano()/1000000), crand.Reader).String())
	// Output: 01F9VX81HX16CWPVPRZ1HY23FC
}

func ExampleXID() {
	fmt.Println(xid.New().String())
	// Output: c3hkhd86n88o0jdjgatg
}

func ExampleNanoID() {
	id, err := nanoid.New()
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
	// Output: rN1FcPY7_rX5lJ-xHl0LW
}

func ExampleKSUID() {
	fmt.Println(ksuid.New().String())
	// Output: 1uv0Ma3iQrWyQGmWjMYLDEik8BX
}

func ExampleSandflake() {
	var g sandflake.Generator

	fmt.Println(g.Next().String())
	// Output: 05X7JEHZ190PR7A70000133R80
}

func ExampleSnowflake() {
	n, _ := snowflake.NewNode(255)

	fmt.Println(n.Generate().String())
	// Output: 1412207541694230528
}

func ExampleSonyflake() {
	var st sonyflake.Settings
	st.StartTime = time.Now()
	sf := sonyflake.NewSonyflake(st)

	id, err := sf.NextID()
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
	// Output: 16779884
}

func ExampleShortUUID() {
	fmt.Println(shortuuid.New())
	// Output: AaUfaFQh9GxnsHBHKbjiLE
}

func ExampleHashID() {
	hdata := hashid.NewData()
	hdata.MinLength = 8
	hdata.Salt = "my salt"
	hid, _ := hashid.NewWithData(hdata)

	id, err := hid.Encode([]int{1})
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
	// Output: On5OLgYy
}

// docker run --rm -d --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=pass -e MYSQL_DATABASE=test mysql:latest
func ExampleAutoIncrement() {
	type Resource struct {
		ID int `gorm:"type:int AUTO_INCREMENT"`
	}
	db, err := gorm.Open(mysql.Open("root:pass@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(Resource{})

	r := Resource{}
	db.Create(&r)

	fmt.Println(r.ID)
	// Output: 1
}

// docker run --rm -d --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=pass -e MYSQL_DATABASE=test mysql:latest
func ExampleAutoIncrementWithHashID() {
	type Resource struct {
		ID int `gorm:"type:int AUTO_INCREMENT"`
	}
	db, err := gorm.Open(mysql.Open("root:pass@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(Resource{})

	hdata := hashid.NewData()
	hdata.MinLength = 8
	hdata.Salt = "my salt"
	hid, _ := hashid.NewWithData(hdata)

	r := Resource{}
	_ = db.Create(&r)

	id, err := hid.Encode([]int{r.ID})
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
	// Output: XPjowja0
}
