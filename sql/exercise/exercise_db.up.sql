-- Define custom enumeration type for muscle groups
CREATE TYPE muscle_group_enum AS ENUM ('arms', 'legs', 'chest', 'back', 'shoulders', 'core');

-- Create exercise table
CREATE TABLE exercise (
    id SERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NULL,
    muscle_group muscle_group_enum NOT NULL,
    user_id INTEGER NOT NULL
);

-- Create equipment table
CREATE TABLE equipment (
    id BIGINT NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    user_id INTEGER NOT NULL
);

-- Create exercise_equipment table
CREATE TABLE exercise_equipment (
    equipment_id BIGINT NOT NULL,
    exercise_id BIGINT NOT NULL,
    PRIMARY KEY (equipment_id, exercise_id),
    CONSTRAINT exercise_equipment_exercise_id_fk FOREIGN KEY (exercise_id) REFERENCES exercise (id),
    CONSTRAINT exercise_equipment_equipment_id_fk FOREIGN KEY (equipment_id) REFERENCES equipment (id)
);
