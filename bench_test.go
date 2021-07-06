package genidbench

import (
	crand "crypto/rand"
	"testing"
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

func benchmark(b *testing.B, f func()) {
	b.ReportAllocs()

	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		f()
	}
}

func BenchmarkUUIDv1(b *testing.B) {
	benchmark(b, func() {
		_ = google_uuid.Must(google_uuid.NewUUID()).String()
	})
}
func BenchmarkUUIDv3(b *testing.B) {
	benchmark(b, func() {
		_ = gofrs_uuid.NewV3(gofrs_uuid.NamespaceDNS, "example.com").String()
	})
}

func BenchmarkUUIDv4(b *testing.B) {
	benchmark(b, func() {
		_ = google_uuid.New().String()
	})
}

func BenchmarkUUIDv5(b *testing.B) {
	benchmark(b, func() {
		_ = gofrs_uuid.NewV5(gofrs_uuid.NamespaceDNS, "example.com").String()
	})
}

func BenchmarkUUIDv6(b *testing.B) {
	benchmark(b, func() {
		_ = uuiddraft.Must(uuiddraft.NewV6()).String()
	})
}

func BenchmarkUUIDv7(b *testing.B) {
	benchmark(b, func() {
		_ = uuiddraft.Must(uuiddraft.NewV7()).String()
	})
}

func BenchmarkULID(b *testing.B) {
	benchmark(b, func() {
		_ = ulid.MustNew(uint64(time.Now().UnixNano()/1000000), crand.Reader).String()
	})
}

func BenchmarkXID(b *testing.B) {
	benchmark(b, func() {
		_ = xid.New().String()
	})
}

func BenchmarkNanoID(b *testing.B) {
	benchmark(b, func() {
		_, err := nanoid.New()
		if err != nil {
			panic(err)
		}
	})
}

func BenchmarkKSUID(b *testing.B) {
	benchmark(b, func() {
		_ = ksuid.New().String()
	})
}

func BenchmarkSandflake(b *testing.B) {
	var g sandflake.Generator
	benchmark(b, func() {
		_ = g.Next().String()
	})
}

func BenchmarkSnowflake(b *testing.B) {
	n, _ := snowflake.NewNode(255)
	benchmark(b, func() {

		_ = n.Generate().String()
	})
}

func BenchmarkSonyflake(b *testing.B) {
	var st sonyflake.Settings
	st.StartTime = time.Now()
	sf := sonyflake.NewSonyflake(st)

	benchmark(b, func() {
		_, err := sf.NextID()
		if err != nil {
			panic(err)
		}
	})
}

func BenchmarkShortUUID(b *testing.B) {
	benchmark(b, func() {
		_ = shortuuid.New()
	})
}

func BenchmarkHashID(b *testing.B) {
	hid, _ := hashid.New()

	benchmark(b, func() {
		_, err := hid.Encode([]int{1})
		if err != nil {
			panic(err)
		}
	})
}

// docker run --rm -d --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=pass -e MYSQL_DATABASE=test mysql:latest
func BenchmarkAutoIncrement(b *testing.B) {
	type Resource struct {
		ID int `gorm:"type:int AUTO_INCREMENT"`
	}
	db, err := gorm.Open(mysql.Open("root:pass@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(Resource{})

	benchmark(b, func() {
		r := Resource{}
		_ = db.Create(&r)
		_ = r.ID
	})
}

// docker run --rm -d --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=pass -e MYSQL_DATABASE=test mysql:latest
func BenchmarkAutoIncrementWithHashID(b *testing.B) {
	type Resource struct {
		ID int `gorm:"type:int AUTO_INCREMENT"`
	}
	db, err := gorm.Open(mysql.Open("root:pass@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(Resource{})

	hid, _ := hashid.New()

	benchmark(b, func() {
		r := Resource{}
		_ = db.Create(&r)

		_, err := hid.Encode([]int{r.ID})
		if err != nil {
			panic(err)
		}
	})
}
