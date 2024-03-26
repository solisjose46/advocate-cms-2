CREATE TABLE Authors (
    author_id INTEGER,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    PRIMARY KEY(author_id)
);

CREATE TABLE Images (
    image_id INTEGER,
    image_src TEXT,
    image_alt TEXT NOT NULL,
    PRIMARY KEY(image_id)
);

CREATE TABLE Stories (
    story_id INTEGER AUTO_INCREMENT,
    file_name TEXT NOT NULL,
    story_name TEXT NOT NULL,
    release_category TEXT NOT NULL,
    content_category TEXT NOT NULL,
    PRIMARY KEY(story_id)
);

CREATE TABLE StoryInfo (
    story_id INTEGER,
    author_id INTEGER NOT NULL,
    description TEXT NOT NULL,
    publish_date DATE NOT NULL,
    PRIMARY KEY(story_id),
    FOREIGN KEY(story_id) REFERENCES Stories(story_id)
    FOREIGN KEY(author_id) REFERENCES Authors(author_id)
);

-- The image with pk story_id, order_num = 0 is the thumbnail

CREATE TABLE StoryImages (
    story_id INTEGER,
    order_num INTEGER NOT NULL,
    image_id INTEGER NOT NULL,
    caption TEXT NOT NULL,
    PRIMARY KEY(story_id, order_num),
    FOREIGN KEY(story_id) REFERENCES Stories(story_id),
    FOREIGN KEY(image_id) REFERENCES Images(image_id)
);

CREATE TABLE StoryParagraphs (
    story_id INTEGER,
    order_num INTEGER NOT NULL,
    paragraph TEXT NOT NULL,
    PRIMARY KEY(story_id, order_num),
    FOREIGN KEY(story_id) REFERENCES Stories(story_id)
);