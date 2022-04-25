CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    authors JSONb NOT NULL,
    subarticles JSONb NOT NULL
);

INSERT INTO
    articles(id, authors, subarticles)
values
    (
        1,
        '[
        {
            "name": "alice",
            "password": "azerty",
            "groups":         ["Redactors", "Users"],
            "email":          "alice@example.com",
            "commonName":     "Alice Smith",
            "surname":        "Smith",
            "givenName":      "Alice"
        },
        {
            "name": "bob",
            "password": "azerty",
            "groups":         ["Redactors", "Users"],
            "email":          "bob@example.com",
            "commonName":     "Bob Smith",
            "surname":        "Smith",
            "givenName":      "Bob"
        },
        {
            "name": "charlie",
            "password": "azerty",
            "groups":         ["Redactors", "Users"],
            "email":          "charlie@example.com",
            "commonName":     "Charlie Smith",
            "surname":        "Smith",
            "givenName":      "Charlie"
        }
    ]',
        '[
        {
            "id":        "1",
            "published":    true,
            "authors":  ["alice", "bob"],
            "categories": ["category1", "category2"],
            "tags":   ["tech"]
        },
        {
            "id":        "2",
            "published":    false,
            "authors":  ["bob"],
            "categories": ["category0"],
            "tags":   ["tech"]
        },
        {
            "id":        "3",
            "published":    false,
            "authors":  ["bob"],
            "categories": ["category3"],
            "tags":   ["sport"]
        },
        {
            "id":        "4",
            "published":    true,
            "authors":  ["charlie"],
            "categories": ["category3"],
            "tags":   ["wholesome"]
        }
    ]'
    );

select
    items.published
from
    articles,
    jsonb_to_recordset(subarticles) as items(id text, published bool)
where
    items.id = '2';


select
    items.id, items.published, items.authors, items.categories, items.tags
from
    articles,
    jsonb_to_recordset(subarticles) as items(id text, published bool, authors text, categories text, tags text)
where
    items.id = '2';
    