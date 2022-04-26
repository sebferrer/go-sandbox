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

SELECT
    *
FROM
    articles
WHERE
    id = 1;

select
    items.published
from
    articles,
    jsonb_to_recordset(subarticles) as items(id text, published bool)
where
    items.id = '2';

select
    items.id,
    items.published,
    items.authors,
    items.categories,
    items.tags
from
    articles,
    jsonb_to_recordset(subarticles) as items(
        id text,
        published bool,
        authors text,
        categories text,
        tags text
    )
where
    items.id = '2';

UPDATE
    articles
SET
    subarticles = subarticles || '{"id":"5","published":true,"authors":["charlie"],"categories":["category3"],"tags":["cats"]}' :: jsonb
WHERE
    id = 1;

UPDATE
    articles i
SET
    subarticles = i2.subarticles
FROM
    (
        SELECT
            id,
            array_to_json(array_agg(elem)) AS subarticles
        FROM
            articles i2,
            json_array_elements(i2.subarticles :: json) elem
        WHERE
            elem ->> 'id' <> '5'
        GROUP BY
            1
    ) i2
WHERE
    i2.id = i.id
    AND json_array_length(i2.subarticles) < json_array_length(i.subarticles :: json);

