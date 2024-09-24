CREATE TABLE repositories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    org_user TEXT NOT NULL,
    repo TEXT NOT NULL,
    created_at int,
    download_url TEXT
);

CREATE UNIQUE INDEX repositories_unique_repo ON repositories (org_user, repo);
