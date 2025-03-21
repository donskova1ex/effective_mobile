-- +goose Up
-- +goose StatementBegin
CREATE TABLE songs (
    id SERIAL PRIMARY KEY,
    group_name VARCHAR(255) NOT NULL,
    song_name VARCHAR(255) NOT NULL,
    release_date DATE NOT NULL,
    text TEXT NOT NULL,
    link VARCHAR(255) NOT NULL
);
CREATE UNIQUE INDEX idx_group_song ON songs (group_name, song_name);

-- Test data
INSERT INTO songs (group_name, song_name, release_date, text, link)
VALUES 
    (
        'The Beatles', 
        'Hey Jude', 
        '1968-08-26', 
        'Hey Jude, don''t make it bad. Take a sad song and make it better.', 
        'https://example.com/the-beatles-hey-jude'
    ),
    (
        'Queen', 
        'Bohemian Rhapsody', 
        '1975-10-31', 
        'Is this the real life? Is this just fantasy?', 
        'https://example.com/queen-bohemian-rhapsody'
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS songs;
-- +goose StatementEnd
