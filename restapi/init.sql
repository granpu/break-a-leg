DROP TABLE IF EXISTS tasks;

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    description TEXT,
    completed BOOLEAN DEFAULT false,
    create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);

INSERT INTO tasks (title, description, completed) VALUES
(title 'изучить GO', description 'пройти базовый курс', completed 'true'),
(title 'решить REST API', description 'написать самому', completed 'false'),
(title 'сделать приложение', description 'ыыыыыххх', completed 'false');

SELECT * FROM tasks;
