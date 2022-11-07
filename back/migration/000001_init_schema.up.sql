CREATE TABLE "answers" (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_name STRING NOT NULL,
    question_id STRING NOT NULL,
    answer STRING NOT NULL
);

CREATE TABLE "questions" (
    id STRING PRIMARY KEY NOT NULL,
    context STRING NOT NULL,
    question STRING NOT NULL
);