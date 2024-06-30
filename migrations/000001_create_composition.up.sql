CREATE TYPE composition_status AS ENUM ('draft', 'in_progress', 'completed', 'published');

CREATE TABLE compositions (
                              composition_id uuid default gen_random_uuid() PRIMARY KEY,
                              user_id uuid NOT NULL,
                              title VARCHAR(100) NOT NULL,
                              description TEXT,
                              status composition_status DEFAULT 'draft',
                              created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                              deleted_at BIGINT default 0
);

CREATE TABLE tracks (
                        track_id uuid default gen_random_uuid() PRIMARY KEY,
                        composition_id uuid REFERENCES compositions(composition_id),
                        user_id uuid NOT NULL,
                        title VARCHAR(100) NOT NULL,
                        file_url VARCHAR(255) NOT NULL,
                        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                        deleted_at BIGINT default 0
);
