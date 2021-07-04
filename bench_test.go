package genidbench

import (
	crand "crypto/rand"
	"testing"
	"time"

	uuiddraft "github.com/coding-socks/uuiddraft"
	gofrs_uuid "github.com/gofrs/uuid"
	google_uuid "github.com/google/uuid"
	shortuuid "github.com/lithammer/shortuuid/v3"
	ulid "github.com/oklog/ulid/v2"
	xid "github.com/rs/xid"
	snowflake "github.com/sony/sonyflake"
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

func BenchmarkUUIDv1NewString(b *testing.B) {
	benchmark(b, func() {
		_ = google_uuid.Must(google_uuid.NewUUID()).String()
	})
}
func BenchmarkUUIDv3NewString(b *testing.B) {
	benchmark(b, func() {
		_ = gofrs_uuid.NewV3(gofrs_uuid.NamespaceDNS, "example.com").String()
	})
}

func BenchmarkUUIDv4NewString(b *testing.B) {
	benchmark(b, func() {
		_ = google_uuid.New().String()
	})
}

func BenchmarkUUIDv5NewString(b *testing.B) {
	benchmark(b, func() {
		_ = gofrs_uuid.NewV5(gofrs_uuid.NamespaceDNS, "example.com").String()
	})
}

func BenchmarkUUIDv6NewString(b *testing.B) {
	benchmark(b, func() {
		_ = uuiddraft.Must(uuiddraft.NewV6()).String()
	})
}

func BenchmarkUUIDv7NewString(b *testing.B) {
	benchmark(b, func() {
		_ = uuiddraft.Must(uuiddraft.NewV7()).String()
	})
}

func BenchmarkSnowFlakeNewUint64(b *testing.B) {
	var st snowflake.Settings
	st.StartTime = time.Now()
	sf := snowflake.NewSonyflake(st)

	benchmark(b, func() {
		_, err := sf.NextID()
		if err != nil {
			panic(err)
		}
	})
}

func BenchmarkULIDNewString(b *testing.B) {
	benchmark(b, func() {
		_ = ulid.MustNew(uint64(time.Now().UnixNano()/1000000), crand.Reader).String()
	})
}

func BenchmarkXIDNewString(b *testing.B) {
	benchmark(b, func() {
		_ = xid.New().String()
	})
}

func BenchmarkShortUUIDNewString(b *testing.B) {
	benchmark(b, func() {
		_ = shortuuid.New()
	})
}

func BenchmarkHashIDEncodeString(b *testing.B) {
	hid, _ := hashid.New()

	benchmark(b, func() {
		_, err := hid.Encode([]int{1})
		if err != nil {
			panic(err)
		}
	})
}

// docker run --rm -d --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=pass -e MYSQL_DATABASE=test mysql:latest
func BenchmarkAutoIncrementString(b *testing.B) {
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
func BenchmarkAutoIncrementWithHashIDEncodeString(b *testing.B) {
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
