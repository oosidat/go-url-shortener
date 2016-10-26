package stores

import(
  "time"
  "github.com/gocql/gocql"
  "github.com/pkg/errors"
)

const (
  ShortURLTable = "short_urls"
)

type Cassandra struct {
	Session *gocql.Session
}

type Config struct {
	Cluster  []string
	Keyspace string
	User     string
	Password string
	Timeout  time.Duration
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

	if cfg.Timeout > 0 {
		c.Timeout = cfg.Timeout
	}

	session, err := c.CreateSession()
	if err != nil {
		return errors.Wrapf(err, "could not create cassandra session")
	}
  s.Session = session
	return err
}

func (s *Cassandra) Code() string {
	return gocql.TimeUUID().String()
}

func (s *Cassandra) Save(url string) (string, error) {
  code := s.Code()
  insertQuery := "INSERT INTO example.short_urls (id, long_url) VALUES (?, ?)"
  err := s.Session.Query(insertQuery, code, url).Exec()
  if err != nil {
    return "", err
  }
  return code, err

}

func (s *Cassandra) Load(code string) (string, error) {
  var longURL string
  selectQuery := "SELECT long_url FROM example.short_urls WHERE id = ? LIMIT 1"
  err := s.Session.Query(selectQuery, code).Consistency(gocql.One).Scan(&longURL);
  if err != nil {
    return "", err
  }
  return longURL, err
}

func (s *Cassandra) Close() {
  s.Session.Close()
}
