CREATE TABLE "StatusChanges" (
                                 "issueId" int,
                                 "authorId" int,
                                 "changeTime" timestamp,
                                 "fromStatus" text,
                                 "toStatus" text
);

CREATE TABLE "Issue" (
                         "id" SERIAL PRIMARY KEY,
                         "projectId" int,
                         "authorId" int,
                         "assigneeId" int,
                         "key" text,
                         "summary" text,
                         "description" text,
                         "type" text,
                         "priority" text,
                         "status" text,
                         "createdTime" timestamp,
                         "closedTime" timestamp,
                         "updatedTime" timestamp,
                         "timeSpent" int
);

CREATE TABLE "Project" (
                           "id" SERIAL PRIMARY KEY,
                           "title" text
);

CREATE TABLE "Author" (
                          "id" SERIAL PRIMARY KEY,
                          "name" text
);

ALTER TABLE "StatusChanges" ADD FOREIGN KEY ("authorId") REFERENCES "Author" ("id");

ALTER TABLE "Issue" ADD FOREIGN KEY ("assigneeId") REFERENCES "Author" ("id");

ALTER TABLE "Issue" ADD FOREIGN KEY ("authorId") REFERENCES "Author" ("id");

ALTER TABLE "StatusChanges" ADD FOREIGN KEY ("issueId") REFERENCES "Issue" ("id");

ALTER TABLE "Issue" ADD FOREIGN KEY ("projectId") REFERENCES "Project" ("id");