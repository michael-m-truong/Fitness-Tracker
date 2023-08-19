-- Create workout_exercise table
CREATE TABLE "workout_exercise" (
    "workout_id" BIGINT NOT NULL,
    "exercise_id" BIGINT NOT NULL,
    "duration" INTEGER NULL,
    "sets" INTEGER NULL,
    "reps" INTEGER NULL,
    "weight" INTEGER NULL,
    PRIMARY KEY ("workout_id", "exercise_id"),
    CONSTRAINT "fk_workout_exercise_workout_id" FOREIGN KEY ("workout_id") REFERENCES "workout" ("id"),
    CONSTRAINT "fk_workout_exercise_exercise_id" FOREIGN KEY ("exercise_id") REFERENCES "exercise" ("id")
);

-- Create workout table
CREATE TABLE "workout" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "user_id" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL
);
