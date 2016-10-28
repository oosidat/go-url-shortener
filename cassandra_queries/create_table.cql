CREATE TABLE example.short_urls (
  id text,
  long_url text,
  context map<text, text>,
  PRIMARY KEY (id, long_url)
);
