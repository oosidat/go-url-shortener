package stores

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/oosidat/go-url-shortener/app"
	"github.com/pkg/errors"
	"github.com/ventu-io/go-shortid"
	"time"
)

const (
	ShortURLTable = "short_urls"
)

type Cassandra struct {
	Session  *gocql.Session
	ShortId  *shortid.Shortid
	Keyspace string
}

type Config struct {
	Cluster  []string
	Keyspace string
	User     string
	Password string
	Timeout  time.Duration
	Seed     uint64
	Worker   uint8
}

func (s *Cassandra) Init(cfg Config) error {
	c := gocql.NewCluster(cfg.Cluster...)
	if cfg.User != "" && cfg.Password != "" {
		c.Authenticator = gocql.PasswordAuthenticator{
			Username: cfg.User,
			Password: cfg.Password,
		}
	}
	c.ProtoVersion = 4
	c.Keyspace = cfg.Keyspace
	s.Keyspace = cfg.Keyspace

	if cfg.Timeout > 0 {
		c.Timeout = cfg.Timeout
	}

	session, err := c.CreateSession()
	if err != nil {
		return errors.Wrapf(err, "could not create cassandra session")
	}
	s.Session = session

	sid, err := shortid.New(cfg.Worker, shortid.DefaultABC, cfg.Seed)
	if err != nil {
		return errors.Wrapf(err, "could not create short id generator")
	}
	s.ShortId = sid

	return err
}

func (s *Cassandra) Code() (string, error) {
	code, err := s.ShortId.Generate()
	if err != nil {
		return "", err
	}
	return code, err
}

func (s *Cassandra) Save(url string) (string, error) {
	code, codeGenerationErr := s.Code()
	if codeGenerationErr != nil {
		return "", codeGenerationErr
	}

	insertQuery := fmt.Sprintf("INSERT INTO %s.%s (id, long_url) VALUES (?, ?)", s.Keyspace, ShortURLTable)
	err := s.Session.Query(insertQuery, code, url).Exec()
	if err != nil {
		return "", err
	}
	return code, err
}

func (s *Cassandra) SavePayload(payload *app.ShortURLCreatePayload) (string, error) {
	code, codeGenerationErr := s.Code()
	if codeGenerationErr != nil {
		return "", codeGenerationErr
	}

	insertQuery := fmt.Sprintf("INSERT INTO %s.%s (id, long_url, context) VALUES (?, ?, ?)", s.Keyspace, ShortURLTable)
	err := s.Session.Query(insertQuery, code, payload.URL, payload.Context).Exec()
	if err != nil {
		return "", err
	}
	return code, err
}

func (s *Cassandra) Load(code string) (string, error) {
	var longURL string
	selectQuery := fmt.Sprintf("SELECT long_url FROM %s.%s WHERE id = ? LIMIT 1", s.Keyspace, ShortURLTable)
	err := s.Session.Query(selectQuery, code).Consistency(gocql.One).Scan(&longURL)
	if err != nil {
		return "", err
	}
	return longURL, err
}

func (s *Cassandra) LoadRecord(code string) (app.GoaExampleShortURL, error) {
	var longURL string
	var context map[string]string

	selectQuery := fmt.Sprintf("SELECT long_url, context FROM %s.%s WHERE id = ? LIMIT 1", s.Keyspace, ShortURLTable)
	err := s.Session.Query(selectQuery, code).Consistency(gocql.One).Scan(&longURL, &context)
	if err != nil {
		return app.GoaExampleShortURL{}, err
	}

	return app.GoaExampleShortURL{LongURL: &longURL, ShortURL: code, Context: context}, err

}

func (s *Cassandra) Close() {
	s.Session.Close()
}
