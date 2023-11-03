create table public.author (
                               id uuid primary key default gen_random_uuid(),
                               name varchar(100) not null
);

CREATE TABLE book (
                      id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                      name VARCHAR(100) not null,
                      author_id uuid not null,
                      constraint author_fk foreign key (author_id) references  public.author(id)
);

insert into author (name) values('Народ');
insert into author (name) values('Джоан Роулинг');
insert into author (name) values('Серега Есенин');

-- insert into book (name, author_id) values ('колобок','47cae5db-b1c7-4920-bf82-5f35b13cf1bb');
-- insert into book (name, author_id) values ('Гарри Поттер','0f6ca29b-8004-446a-a74b-1633d1a4dc94');
-- insert into book (name, author_id) values ('Пой же, Пой','dbe28e77-26cf-4409-9e46-b563186f155b');