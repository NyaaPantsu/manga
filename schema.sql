--
-- PostgreSQL database dump
--

-- Dumped from database version 10.1
-- Dumped by pg_dump version 10.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: batoto; Type: SCHEMA; Schema: -; Owner: manga
--

CREATE SCHEMA batoto;


ALTER SCHEMA batoto OWNER TO manga;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

--
-- Name: language; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE language AS ENUM (
    'English',
    'Spanish',
    'French',
    'German',
    'Portuguese',
    'Turkish',
    'Indonesian',
    'Greek',
    'Filipino',
    'Italian',
    'Polish',
    'Thai',
    'Malay',
    'Hungarian',
    'Romanian',
    'Arabic',
    'Hebrew',
    'Russian',
    'Vietnamese',
    'Dutch',
    'Bengali',
    'Persian',
    'Czech',
    'Brazilian',
    'Bulgarian',
    'Danish',
    'Esperanto',
    'Swedish',
    'Lithuanian',
    'Other'
);


ALTER TYPE language OWNER TO postgres;

--
-- Name: status; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE status AS ENUM (
    'Stalled',
    'Discontinued',
    'Ongoing',
    'Complete'
);


ALTER TYPE status OWNER TO postgres;

--
-- Name: type; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE type AS ENUM (
    'Manga (Japanese)',
    'Manhwa (Korean)',
    'Webcomic (Japanese)',
    'Webcomic (Korean)'
);


ALTER TYPE type OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: groups_scanlation; Type: TABLE; Schema: public; Owner: manga
--

CREATE TABLE groups_scanlation (
    name text NOT NULL,
    description text NOT NULL,
    release_delay integer
);


ALTER TABLE groups_scanlation OWNER TO manga;

--
-- Name: TABLE groups_scanlation; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON TABLE groups_scanlation IS 'Describes scanlation groups; groups who scan and translate series.';


--
-- Name: COLUMN groups_scanlation.name; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN groups_scanlation.name IS 'Name to refer to the scanlation group by.';


--
-- Name: COLUMN groups_scanlation.description; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN groups_scanlation.description IS 'Human readable description for the scanlation group.';


--
-- Name: COLUMN groups_scanlation.release_delay; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN groups_scanlation.release_delay IS 'A fixed delay, counted from the time a chapter is uploaded, before the chapter should be released.';


--
-- Name: groups_scanlation_urls; Type: TABLE; Schema: public; Owner: manga
--

CREATE TABLE groups_scanlation_urls (
    group_name text NOT NULL,
    url text NOT NULL
);


ALTER TABLE groups_scanlation_urls OWNER TO manga;

--
-- Name: TABLE groups_scanlation_urls; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON TABLE groups_scanlation_urls IS 'URLs belonging to a scanlation group.  Might be scanlation group websites, social media pages, etc.';


--
-- Name: COLUMN groups_scanlation_urls.group_name; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN groups_scanlation_urls.group_name IS 'Identifies the scanlation group the link belongs to.\n\nHas foreign key relationship with the table of scanlation groups.';


--
-- Name: COLUMN groups_scanlation_urls.url; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN groups_scanlation_urls.url IS 'The actual URL associated with the scanlation group.';


--
-- Name: languages; Type: TABLE; Schema: public; Owner: manga
--

CREATE TABLE languages (
    name text NOT NULL,
    code text
);


ALTER TABLE languages OWNER TO manga;

--
-- Name: TABLE languages; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON TABLE languages IS 'Describes acceptable languages for series chapters.';


--
-- Name: COLUMN languages.name; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN languages.name IS 'Contains the name of the language, e.g. "English".';


--
-- Name: series; Type: TABLE; Schema: public; Owner: manga
--

CREATE TABLE series (
    id bigint NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    cover_image text NOT NULL,
    type_name text NOT NULL,
    type_demonym text NOT NULL,
    status text NOT NULL
);


ALTER TABLE series OWNER TO manga;

--
-- Name: TABLE series; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON TABLE series IS 'Describes manga, manhwa, etc series.';


--
-- Name: COLUMN series.id; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series.id IS 'Unique automatically generated sequential identifier for the series.';


--
-- Name: COLUMN series.name; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series.name IS 'Preferred name to refer to a series by.';


--
-- Name: COLUMN series.description; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series.description IS 'Human readable description of the series.';


--
-- Name: COLUMN series.cover_image; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series.cover_image IS 'IPFS hash for the preferred cover image for the series.';


--
-- Name: COLUMN series.type_name; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series.type_name IS 'Name of the series type; e.g. "Manga" or "Manhwa" or "Webcomic".\n\nHas a foreign key relationship with the table of allowable series types.';


--
-- Name: COLUMN series.type_demonym; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series.type_demonym IS 'Demonym of the series type; e.g. "Japanese" or "Korean".\n\nHas foreign key relationship with the table of allowable series types.';


--
-- Name: COLUMN series.status; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series.status IS 'Status of the series, such as "Complete" or "Stalled".\n\nHas foreign key relationship with the table of allowable statuses.';


--
-- Name: series_aliases; Type: TABLE; Schema: public; Owner: manga
--

CREATE TABLE series_aliases (
    series_id bigint NOT NULL,
    name text NOT NULL
);


ALTER TABLE series_aliases OWNER TO manga;

--
-- Name: TABLE series_aliases; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON TABLE series_aliases IS 'Contains the alternative names for series.';


--
-- Name: COLUMN series_aliases.series_id; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series_aliases.series_id IS 'Identifies the series the alternative name belongs to.\n\nHas foreign key relationship with the table of series.';


--
-- Name: COLUMN series_aliases.name; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series_aliases.name IS 'Alternative name to refer to the series by.';


--
-- Name: series_chapters; Type: TABLE; Schema: public; Owner: manga
--

CREATE TABLE series_chapters (
    id bigint NOT NULL,
    series_id bigint NOT NULL,
    title text NOT NULL,
    chapter_number_absolute text NOT NULL,
    chapter_number_volume numeric(10,2),
    volume_number numeric(10,2),
    chapter_language text NOT NULL,
    contributor_id bigint,
    time_uploaded timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    hash text NOT NULL
);


ALTER TABLE series_chapters OWNER TO manga;

--
-- Name: TABLE series_chapters; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON TABLE series_chapters IS 'Describes chapters of a series.';


--
-- Name: COLUMN series_chapters.id; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series_chapters.id IS 'Uniquely identifies the series chapter.\n\nID column is used here rather than a primary key encapsulating the series ID and chapter number because the support for series volumes would make this primary key unduly complicated.';


--
-- Name: COLUMN series_chapters.series_id; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series_chapters.series_id IS 'Identifies the series the chapter belongs to.\n\nHas foreign key relationship with the series table.';


--
-- Name: COLUMN series_chapters.title; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series_chapters.title IS 'Title of the chapter.';


--
-- Name: COLUMN series_chapters.chapter_number_absolute; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series_chapters.chapter_number_absolute IS 'Absolute chapter number; the chapter number relative to the total number of chapters in a series, NOT relative to a volume.';


--
-- Name: COLUMN series_chapters.chapter_number_volume; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series_chapters.chapter_number_volume IS 'Chapter number relative to the number of chapters in a volume.  This is the number of the chapter WITHIN the a volume.';


--
-- Name: COLUMN series_chapters.volume_number; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series_chapters.volume_number IS 'The number of the volume the chapter belongs to.';


--
-- Name: COLUMN series_chapters.chapter_language; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series_chapters.chapter_language IS 'Language of the text content of the chapter.\n\nHas foreign key relationship with the languages table.';


--
-- Name: COLUMN series_chapters.contributor_id; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series_chapters.contributor_id IS 'Identifies the user who contributed the chapter.\n\nHas foreign key relationship with the table of users.';


--
-- Name: COLUMN series_chapters.time_uploaded; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series_chapters.time_uploaded IS 'Timestamp of when the chapter was uploaded.\n\nDefaults to current time.';


--
-- Name: COLUMN series_chapters.hash; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series_chapters.hash IS 'IPFS hash of the folder for the series chapter.';


--
-- Name: series_chapters_files; Type: TABLE; Schema: public; Owner: manga
--

CREATE TABLE series_chapters_files (
    chapter_id bigint NOT NULL,
    name text NOT NULL
);


ALTER TABLE series_chapters_files OWNER TO manga;

--
-- Name: TABLE series_chapters_files; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON TABLE series_chapters_files IS 'Describes the files in a chapter''s IPFS folder.\n\nTODO: Someone please justify this table''s existence, as its necessity is not apparent to the schema author.';


--
-- Name: COLUMN series_chapters_files.chapter_id; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series_chapters_files.chapter_id IS 'Identifier for the chapter the file relates to.\n\nHas foreign key relationship with the series_chapters table.';


--
-- Name: COLUMN series_chapters_files.name; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series_chapters_files.name IS 'Semantic name of the file in the chapter''s folder.';


--
-- Name: series_chapters_groups; Type: TABLE; Schema: public; Owner: manga
--

CREATE TABLE series_chapters_groups (
    chapter_id bigint NOT NULL,
    group_name text NOT NULL
);


ALTER TABLE series_chapters_groups OWNER TO manga;

--
-- Name: TABLE series_chapters_groups; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON TABLE series_chapters_groups IS 'Identifies the scanlation groups a series chapter belongs to.';


--
-- Name: COLUMN series_chapters_groups.chapter_id; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series_chapters_groups.chapter_id IS 'Identifies the chapter that belongs to a given scanlation group.';


--
-- Name: COLUMN series_chapters_groups.group_name; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series_chapters_groups.group_name IS 'Identifies the name of the scanlation group that the chapter belongs to.';


--
-- Name: series_chapters_id_seq; Type: SEQUENCE; Schema: public; Owner: manga
--

CREATE SEQUENCE series_chapters_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE series_chapters_id_seq OWNER TO manga;

--
-- Name: series_chapters_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: manga
--

ALTER SEQUENCE series_chapters_id_seq OWNED BY series_chapters.id;


--
-- Name: series_id_seq; Type: SEQUENCE; Schema: public; Owner: manga
--

CREATE SEQUENCE series_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE series_id_seq OWNER TO manga;

--
-- Name: series_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: manga
--

ALTER SEQUENCE series_id_seq OWNED BY series.id;


--
-- Name: series_ratings; Type: TABLE; Schema: public; Owner: manga
--

CREATE TABLE series_ratings (
    series_id bigint NOT NULL,
    user_id bigint NOT NULL,
    rating integer NOT NULL
);


ALTER TABLE series_ratings OWNER TO manga;

--
-- Name: TABLE series_ratings; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON TABLE series_ratings IS 'Describes ratings of a series that users have made.';


--
-- Name: COLUMN series_ratings.series_id; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series_ratings.series_id IS 'Identifies the series that is being rated.';


--
-- Name: COLUMN series_ratings.user_id; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series_ratings.user_id IS 'Identifies the user that is rating the series.';


--
-- Name: COLUMN series_ratings.rating; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series_ratings.rating IS 'The actual rating of the series, on a yet to be determined scale.';


--
-- Name: series_tags; Type: TABLE; Schema: public; Owner: manga
--

CREATE TABLE series_tags (
    series_id bigint NOT NULL,
    tag_name text NOT NULL,
    tag_namespace text NOT NULL
);


ALTER TABLE series_tags OWNER TO manga;

--
-- Name: TABLE series_tags; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON TABLE series_tags IS 'Identifies tags associated with a series.';


--
-- Name: COLUMN series_tags.series_id; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series_tags.series_id IS 'Identifies the series the tag is associated with.\n\nHas foreign key relationship with the table of series.';


--
-- Name: COLUMN series_tags.tag_name; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series_tags.tag_name IS 'Identifies the tag associated with the series.\n\nHas foreign key relationship with the table of tags.';


--
-- Name: COLUMN series_tags.tag_namespace; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN series_tags.tag_namespace IS 'Identifies the namespace of the tag associated withh the series.\n\nHas foreign key relationship with the table of tags.';


--
-- Name: statuses; Type: TABLE; Schema: public; Owner: manga
--

CREATE TABLE statuses (
    name text NOT NULL
);


ALTER TABLE statuses OWNER TO manga;

--
-- Name: TABLE statuses; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON TABLE statuses IS 'Describes acceptable statuses for series.';


--
-- Name: COLUMN statuses.name; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN statuses.name IS 'Contains the name of the status, e.g. "Complete".';


--
-- Name: tags; Type: TABLE; Schema: public; Owner: manga
--

CREATE TABLE tags (
    name text NOT NULL,
    namespace text NOT NULL
);


ALTER TABLE tags OWNER TO manga;

--
-- Name: TABLE tags; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON TABLE tags IS 'Describes the acceptable tags for series.';


--
-- Name: COLUMN tags.name; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN tags.name IS 'Describes the name of the tag.';


--
-- Name: COLUMN tags.namespace; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN tags.namespace IS 'Describes the namespace of the tag.';


--
-- Name: types; Type: TABLE; Schema: public; Owner: manga
--

CREATE TABLE types (
    name text NOT NULL,
    origin_demonym text NOT NULL
);


ALTER TABLE types OWNER TO manga;

--
-- Name: TABLE types; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON TABLE types IS 'Describes the allowable values for series types.  Series types describe both the country of origin for a series, and whether it is a web comic.';


--
-- Name: COLUMN types.name; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN types.name IS 'Describes the name of the type of series, e.g. "Manga" or "Manhwa".';


--
-- Name: COLUMN types.origin_demonym; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN types.origin_demonym IS 'Describes the demonym for the country of origin for this series type, e.g. "Japanese" or "Korean".';


--
-- Name: users; Type: TABLE; Schema: public; Owner: manga
--

CREATE TABLE users (
    id bigint NOT NULL,
    username text NOT NULL,
    email text,
    password_hash bytea NOT NULL
);


ALTER TABLE users OWNER TO manga;

--
-- Name: TABLE users; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON TABLE users IS 'Describes user accounts for the Batoto replacement site.';


--
-- Name: COLUMN users.id; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN users.id IS 'Uniquely identifies the user account.\n\nJustification for an identifier column as opposed to using the username as the primary key is simply because it is considered "good practice" to use identifiers instead.';


--
-- Name: COLUMN users.username; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN users.username IS 'Username used as a login credential for the user, as well as a human readable account identifier.';


--
-- Name: COLUMN users.email; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN users.email IS 'Email address associated with a user account.';


--
-- Name: COLUMN users.password_hash; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN users.password_hash IS 'Hash of the user''s password, used as a login credential.  Hash is expected to be a bcrypt hash.';


--
-- Name: users_following_series; Type: TABLE; Schema: public; Owner: manga
--

CREATE TABLE users_following_series (
    user_id bigint NOT NULL,
    series_id bigint NOT NULL
);


ALTER TABLE users_following_series OWNER TO manga;

--
-- Name: TABLE users_following_series; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON TABLE users_following_series IS 'Describes the series that a user is following.';


--
-- Name: COLUMN users_following_series.user_id; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN users_following_series.user_id IS 'Identifies the user that is following a series.';


--
-- Name: COLUMN users_following_series.series_id; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN users_following_series.series_id IS 'Identifies the series that the user is following.';


--
-- Name: users_groups; Type: TABLE; Schema: public; Owner: manga
--

CREATE TABLE users_groups (
    user_id bigint NOT NULL,
    group_name text NOT NULL
);


ALTER TABLE users_groups OWNER TO manga;

--
-- Name: TABLE users_groups; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON TABLE users_groups IS 'Identifies scanlation groups that a user belongs to.';


--
-- Name: COLUMN users_groups.user_id; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN users_groups.user_id IS 'Identifies the user that belongs to a given group.';


--
-- Name: COLUMN users_groups.group_name; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON COLUMN users_groups.group_name IS 'Identifies the scanlation group that the user belongs to.';


--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: manga
--

CREATE SEQUENCE users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE users_id_seq OWNER TO manga;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: manga
--

ALTER SEQUENCE users_id_seq OWNED BY users.id;


--
-- Name: series id; Type: DEFAULT; Schema: public; Owner: manga
--

ALTER TABLE ONLY series ALTER COLUMN id SET DEFAULT nextval('series_id_seq'::regclass);


--
-- Name: series_chapters id; Type: DEFAULT; Schema: public; Owner: manga
--

ALTER TABLE ONLY series_chapters ALTER COLUMN id SET DEFAULT nextval('series_chapters_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: manga
--

ALTER TABLE ONLY users ALTER COLUMN id SET DEFAULT nextval('users_id_seq'::regclass);


--
-- Name: series_chapters idx_series_chapters_hash; Type: CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY series_chapters
    ADD CONSTRAINT idx_series_chapters_hash UNIQUE (hash);


--
-- Name: CONSTRAINT idx_series_chapters_hash ON series_chapters; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT idx_series_chapters_hash ON series_chapters IS 'Indexes the IPFS hash of a chapter.Justification is to be able to identify an otherwise unknown chapter by its IPFS hash.';


--
-- Name: series idx_series_name; Type: CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY series
    ADD CONSTRAINT idx_series_name UNIQUE (name);


--
-- Name: CONSTRAINT idx_series_name ON series; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT idx_series_name ON series IS 'Index indexing the name of the series.';


--
-- Name: groups_scanlation pk_groups_scanlation_name; Type: CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY groups_scanlation
    ADD CONSTRAINT pk_groups_scanlation_name PRIMARY KEY (name);


--
-- Name: groups_scanlation_urls pk_groups_scanlation_urls; Type: CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY groups_scanlation_urls
    ADD CONSTRAINT pk_groups_scanlation_urls PRIMARY KEY (group_name, url);


--
-- Name: CONSTRAINT pk_groups_scanlation_urls ON groups_scanlation_urls; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT pk_groups_scanlation_urls ON groups_scanlation_urls IS 'Primary key encapsulating the name of the scanlation group that a link is associated with, as well as the link itself.';


--
-- Name: languages pk_languages_name; Type: CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY languages
    ADD CONSTRAINT pk_languages_name PRIMARY KEY (name);


--
-- Name: CONSTRAINT pk_languages_name ON languages; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT pk_languages_name ON languages IS 'Primary key encapsulating the name of the language.';


--
-- Name: series_aliases pk_series_aliases; Type: CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY series_aliases
    ADD CONSTRAINT pk_series_aliases PRIMARY KEY (series_id, name);


--
-- Name: CONSTRAINT pk_series_aliases ON series_aliases; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT pk_series_aliases ON series_aliases IS 'Primary key encapsulating the identifier for the series the alias belongs to, and the alias itself.';


--
-- Name: series_chapters_files pk_series_chapters_file; Type: CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY series_chapters_files
    ADD CONSTRAINT pk_series_chapters_file PRIMARY KEY (chapter_id, name);


--
-- Name: CONSTRAINT pk_series_chapters_file ON series_chapters_files; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT pk_series_chapters_file ON series_chapters_files IS 'Primary key encapsulating the identifier of the chapter the file relates to, and the name of the file itself.';


--
-- Name: series_chapters_groups pk_series_chapters_groups; Type: CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY series_chapters_groups
    ADD CONSTRAINT pk_series_chapters_groups PRIMARY KEY (chapter_id, group_name);


--
-- Name: CONSTRAINT pk_series_chapters_groups ON series_chapters_groups; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT pk_series_chapters_groups ON series_chapters_groups IS 'Primary key encapsulating the series chapter that belongs to a given group, as well as the group that it belongs to.';


--
-- Name: series_chapters pk_series_chapters_id; Type: CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY series_chapters
    ADD CONSTRAINT pk_series_chapters_id PRIMARY KEY (id);


--
-- Name: CONSTRAINT pk_series_chapters_id ON series_chapters; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT pk_series_chapters_id ON series_chapters IS 'Primary key encapsulating the unique identifier for the chapter.';


--
-- Name: series pk_series_id; Type: CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY series
    ADD CONSTRAINT pk_series_id PRIMARY KEY (id);


--
-- Name: CONSTRAINT pk_series_id ON series; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT pk_series_id ON series IS 'Primary key encapsulating the automatically generated unique identifier for the series.';


--
-- Name: series_ratings pk_series_ratings; Type: CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY series_ratings
    ADD CONSTRAINT pk_series_ratings PRIMARY KEY (series_id, user_id);


--
-- Name: CONSTRAINT pk_series_ratings ON series_ratings; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT pk_series_ratings ON series_ratings IS 'Primary key encapsulating the series identifier for the series being rated, as well as the user identifier for the user rating the series.';


--
-- Name: series_tags pk_series_tags; Type: CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY series_tags
    ADD CONSTRAINT pk_series_tags PRIMARY KEY (series_id, tag_name, tag_namespace);


--
-- Name: CONSTRAINT pk_series_tags ON series_tags; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT pk_series_tags ON series_tags IS 'Primary key encapsulating the series identifier, tag name, and tag namespace, of the tag<>series association.';


--
-- Name: statuses pk_statuses_name; Type: CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY statuses
    ADD CONSTRAINT pk_statuses_name PRIMARY KEY (name);


--
-- Name: CONSTRAINT pk_statuses_name ON statuses; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT pk_statuses_name ON statuses IS 'Primary key encapsulating the name of the series status.';


--
-- Name: tags pk_tags_name_namespace; Type: CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY tags
    ADD CONSTRAINT pk_tags_name_namespace PRIMARY KEY (name, namespace);


--
-- Name: CONSTRAINT pk_tags_name_namespace ON tags; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT pk_tags_name_namespace ON tags IS 'Primary key encapsulating the name and namespace of the tag.';


--
-- Name: types pk_types; Type: CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY types
    ADD CONSTRAINT pk_types PRIMARY KEY (name, origin_demonym);


--
-- Name: CONSTRAINT pk_types ON types; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT pk_types ON types IS 'Primary key encapsulating both the name of the series type, and the country of origin.';


--
-- Name: users_following_series pk_users_following_series; Type: CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY users_following_series
    ADD CONSTRAINT pk_users_following_series PRIMARY KEY (user_id, series_id);


--
-- Name: CONSTRAINT pk_users_following_series ON users_following_series; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT pk_users_following_series ON users_following_series IS 'Primary key encapsulating the identifier of the user that is following a series, as well as the identifier of the series that the user is following.';


--
-- Name: users_groups pk_users_groups; Type: CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY users_groups
    ADD CONSTRAINT pk_users_groups PRIMARY KEY (user_id, group_name);


--
-- Name: CONSTRAINT pk_users_groups ON users_groups; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT pk_users_groups ON users_groups IS 'Primary key encapsulating the user identifier and the scanlation group name that the user belongs to.';


--
-- Name: users pk_users_id; Type: CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY users
    ADD CONSTRAINT pk_users_id PRIMARY KEY (id);


--
-- Name: users pk_users_username; Type: CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY users
    ADD CONSTRAINT pk_users_username UNIQUE (username);


--
-- Name: CONSTRAINT pk_users_username ON users; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT pk_users_username ON users IS 'Indexes user accounts'' usernames.Justification is to ensure username uniqueness as well as to enable searching of user accounts by username.';


--
-- Name: idx_groups_scanlation_description; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_groups_scanlation_description ON groups_scanlation USING btree (description);


--
-- Name: INDEX idx_groups_scanlation_description; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON INDEX idx_groups_scanlation_description IS 'Indexes the description of a scanlation group.Justification is to make scanlation groups searchabe by their description.';


--
-- Name: idx_groups_scanlation_links_group_name; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_groups_scanlation_links_group_name ON groups_scanlation_urls USING btree (group_name);


--
-- Name: idx_groups_scanlation_release_delay; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_groups_scanlation_release_delay ON groups_scanlation USING btree (release_delay);


--
-- Name: INDEX idx_groups_scanlation_release_delay; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON INDEX idx_groups_scanlation_release_delay IS 'Indexes the release delay for a scanlation group.Justification is to make scanlation groups searchable/sortable by their release delay.';


--
-- Name: idx_series_aliases_series_id; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_series_aliases_series_id ON series_aliases USING btree (series_id);


--
-- Name: INDEX idx_series_aliases_series_id; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON INDEX idx_series_aliases_series_id IS 'Index indexing the series identifiers for aliases.Justification is to allow searching of aliases by series identifier.';


--
-- Name: idx_series_chapters_chapter_language; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_series_chapters_chapter_language ON series_chapters USING btree (chapter_language);


--
-- Name: idx_series_chapters_chapter_number_absolute; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_series_chapters_chapter_number_absolute ON series_chapters USING btree (chapter_number_absolute);


--
-- Name: INDEX idx_series_chapters_chapter_number_absolute; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON INDEX idx_series_chapters_chapter_number_absolute IS 'Indexes the absolute chapter number of a series.Justification is to sort by chapter numbers.';


--
-- Name: idx_series_chapters_chapter_number_volume; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_series_chapters_chapter_number_volume ON series_chapters USING btree (chapter_number_volume);


--
-- Name: INDEX idx_series_chapters_chapter_number_volume; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON INDEX idx_series_chapters_chapter_number_volume IS 'Indexes the relative (volume) chapter number of a series.Justification is to sort by chapter numbers.';


--
-- Name: idx_series_chapters_contributor_id; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_series_chapters_contributor_id ON series_chapters USING btree (contributor_id);


--
-- Name: INDEX idx_series_chapters_contributor_id; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON INDEX idx_series_chapters_contributor_id IS 'Indexes by chapter''s contributor.Justification is to be able to search all chapters contributed by a given user.';


--
-- Name: idx_series_chapters_files_chapter_id; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_series_chapters_files_chapter_id ON series_chapters_files USING btree (chapter_id);


--
-- Name: idx_series_chapters_groups_chapter_id; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_series_chapters_groups_chapter_id ON series_chapters_groups USING btree (chapter_id);


--
-- Name: idx_series_chapters_groups_group_name; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_series_chapters_groups_group_name ON series_chapters_groups USING btree (group_name);


--
-- Name: idx_series_chapters_series_id; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_series_chapters_series_id ON series_chapters USING btree (series_id);


--
-- Name: idx_series_chapters_time_uploaded; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_series_chapters_time_uploaded ON series_chapters USING btree (time_uploaded);


--
-- Name: INDEX idx_series_chapters_time_uploaded; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON INDEX idx_series_chapters_time_uploaded IS 'Indexes the time a chapter was uploaded.Justification is to search by upload time/date.';


--
-- Name: idx_series_chapters_title; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_series_chapters_title ON series_chapters USING btree (title);


--
-- Name: INDEX idx_series_chapters_title; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON INDEX idx_series_chapters_title IS 'Indexes the title of the chapter.Is not unique because series may share chapter titles.Justification for index is to make series_chapters searchable by chapter title.';


--
-- Name: idx_series_chapters_volume_number; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_series_chapters_volume_number ON series_chapters USING btree (volume_number);


--
-- Name: INDEX idx_series_chapters_volume_number; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON INDEX idx_series_chapters_volume_number IS 'Indexes the volume number of a series.Justification is to sort by volume numbers.';


--
-- Name: idx_series_ratings_series_id; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_series_ratings_series_id ON series_ratings USING btree (series_id);


--
-- Name: INDEX idx_series_ratings_series_id; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON INDEX idx_series_ratings_series_id IS 'Indexes the series identifier the rating relates to.\n\nJustification is to allow searching of ratings by series identifier.';


--
-- Name: idx_series_ratings_user_id; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_series_ratings_user_id ON series_ratings USING btree (user_id);


--
-- Name: INDEX idx_series_ratings_user_id; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON INDEX idx_series_ratings_user_id IS 'Indexes the user identifier for a rating.\n\nJustification is to allow searching of ratings by the user who gave the rating.';


--
-- Name: idx_series_series_type_name; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_series_series_type_name ON series USING btree (type_name);


--
-- Name: INDEX idx_series_series_type_name; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON INDEX idx_series_series_type_name IS 'Index indexing the name of the series type.';


--
-- Name: idx_series_status; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_series_status ON series USING btree (status);


--
-- Name: INDEX idx_series_status; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON INDEX idx_series_status IS 'Index indexing the series status.';


--
-- Name: idx_series_tags_series_id; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_series_tags_series_id ON series_tags USING btree (series_id);


--
-- Name: INDEX idx_series_tags_series_id; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON INDEX idx_series_tags_series_id IS 'Indexes the series identifier for the series the tag is being associated with.\n\nJustification is to allow searching tags for all tags assosicated with a series.';


--
-- Name: idx_series_tags_tag_name; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_series_tags_tag_name ON series_tags USING btree (tag_name);


--
-- Name: INDEX idx_series_tags_tag_name; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON INDEX idx_series_tags_tag_name IS 'Indexes the tag name for the series-tag association.\n\nJustification is to allow enumerating all series-tag associations by a tag name.';


--
-- Name: idx_series_tags_tag_name_namespace; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_series_tags_tag_name_namespace ON series_tags USING btree (tag_name, tag_namespace);


--
-- Name: INDEX idx_series_tags_tag_name_namespace; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON INDEX idx_series_tags_tag_name_namespace IS 'Indexes the tag name and namespace.\n\nJustification is to allow searching for series matching given fully formed tags.';


--
-- Name: idx_series_tags_tag_namespace; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_series_tags_tag_namespace ON series_tags USING btree (tag_namespace);


--
-- Name: INDEX idx_series_tags_tag_namespace; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON INDEX idx_series_tags_tag_namespace IS 'Indexes the namespace of the tag associated with the series.\n\nJustification is to allow enumerating all series tags within a given namespace.';


--
-- Name: idx_series_type_demonym; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_series_type_demonym ON series USING btree (type_demonym);


--
-- Name: INDEX idx_series_type_demonym; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON INDEX idx_series_type_demonym IS 'Index indexing the demonym of the series type.';


--
-- Name: idx_users_following_series_series_id; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_users_following_series_series_id ON users_following_series USING btree (series_id);


--
-- Name: idx_users_following_series_user_id; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_users_following_series_user_id ON users_following_series USING btree (user_id);


--
-- Name: idx_users_groups_group_name; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_users_groups_group_name ON users_groups USING btree (group_name);


--
-- Name: idx_users_groups_user_id; Type: INDEX; Schema: public; Owner: manga
--

CREATE INDEX idx_users_groups_user_id ON users_groups USING btree (user_id);


--
-- Name: groups_scanlation_urls fk_groups_scanlation_links; Type: FK CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY groups_scanlation_urls
    ADD CONSTRAINT fk_groups_scanlation_links FOREIGN KEY (group_name) REFERENCES groups_scanlation(name) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: CONSTRAINT fk_groups_scanlation_links ON groups_scanlation_urls; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT fk_groups_scanlation_links ON groups_scanlation_urls IS 'Foreign key binding a scanlation group URL to an actual entry in the scanlation groups table.';


--
-- Name: series_aliases fk_series_aliases_series; Type: FK CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY series_aliases
    ADD CONSTRAINT fk_series_aliases_series FOREIGN KEY (series_id) REFERENCES series(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: CONSTRAINT fk_series_aliases_series ON series_aliases; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT fk_series_aliases_series ON series_aliases IS 'Foreign key relationship binding the series alias to a particular series identifier.';


--
-- Name: series_chapters_files fk_series_chapters_files; Type: FK CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY series_chapters_files
    ADD CONSTRAINT fk_series_chapters_files FOREIGN KEY (chapter_id) REFERENCES series_chapters(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: CONSTRAINT fk_series_chapters_files ON series_chapters_files; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT fk_series_chapters_files ON series_chapters_files IS 'Foreign key binding the identifier of a chapter the file relates to, to an actual entry in the series chapters table.';


--
-- Name: series_chapters_groups fk_series_chapters_groups_chapter_id; Type: FK CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY series_chapters_groups
    ADD CONSTRAINT fk_series_chapters_groups_chapter_id FOREIGN KEY (chapter_id) REFERENCES series_chapters(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: CONSTRAINT fk_series_chapters_groups_chapter_id ON series_chapters_groups; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT fk_series_chapters_groups_chapter_id ON series_chapters_groups IS 'Foreign key binding the identifier of a chapter that belongs to a given group, to an actual entry in the series chapters table.';


--
-- Name: series_chapters_groups fk_series_chapters_groups_group_name; Type: FK CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY series_chapters_groups
    ADD CONSTRAINT fk_series_chapters_groups_group_name FOREIGN KEY (group_name) REFERENCES groups_scanlation(name) ON UPDATE CASCADE;


--
-- Name: CONSTRAINT fk_series_chapters_groups_group_name ON series_chapters_groups; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT fk_series_chapters_groups_group_name ON series_chapters_groups IS 'Foreign key binding the name of a scanlation group that a series chapter belongs to, to an actual entry in the scanlation groups table.';


--
-- Name: series_chapters fk_series_chapters_languages; Type: FK CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY series_chapters
    ADD CONSTRAINT fk_series_chapters_languages FOREIGN KEY (chapter_language) REFERENCES languages(name) ON UPDATE CASCADE;


--
-- Name: CONSTRAINT fk_series_chapters_languages ON series_chapters; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT fk_series_chapters_languages ON series_chapters IS 'Foreign key relationship binding the language of the series chapter to the languages table.';


--
-- Name: series_chapters fk_series_chapters_series_id; Type: FK CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY series_chapters
    ADD CONSTRAINT fk_series_chapters_series_id FOREIGN KEY (series_id) REFERENCES series(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: CONSTRAINT fk_series_chapters_series_id ON series_chapters; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT fk_series_chapters_series_id ON series_chapters IS 'Foreign key relationship binding the series_id to an entry in the series table.';


--
-- Name: series_chapters fk_series_chapters_users; Type: FK CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY series_chapters
    ADD CONSTRAINT fk_series_chapters_users FOREIGN KEY (contributor_id) REFERENCES users(id) ON UPDATE CASCADE;


--
-- Name: CONSTRAINT fk_series_chapters_users ON series_chapters; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT fk_series_chapters_users ON series_chapters IS 'Foreign key binding the identifier of a chapter''s contributor, to an actual entry in the users table.';


--
-- Name: series_ratings fk_series_ratings_series; Type: FK CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY series_ratings
    ADD CONSTRAINT fk_series_ratings_series FOREIGN KEY (series_id) REFERENCES series(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: CONSTRAINT fk_series_ratings_series ON series_ratings; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT fk_series_ratings_series ON series_ratings IS 'Foreign key binding the series identifier to an actual entry in the series table. \n\nCascade on update and delete is to prune entries relating to removed series.';


--
-- Name: series_ratings fk_series_ratings_users; Type: FK CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY series_ratings
    ADD CONSTRAINT fk_series_ratings_users FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: CONSTRAINT fk_series_ratings_users ON series_ratings; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT fk_series_ratings_users ON series_ratings IS 'Foreign key binding the user identifier to an actual entry in the users table. \n\nCascade on delete and update to prune ratings from removed users.';


--
-- Name: series fk_series_statuses; Type: FK CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY series
    ADD CONSTRAINT fk_series_statuses FOREIGN KEY (status) REFERENCES statuses(name) ON UPDATE CASCADE;


--
-- Name: CONSTRAINT fk_series_statuses ON series; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT fk_series_statuses ON series IS 'Foreign key relationship binding series status to the table of allowable series statuses.';


--
-- Name: series_tags fk_series_tags_series; Type: FK CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY series_tags
    ADD CONSTRAINT fk_series_tags_series FOREIGN KEY (series_id) REFERENCES series(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: CONSTRAINT fk_series_tags_series ON series_tags; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT fk_series_tags_series ON series_tags IS 'Foreign key relationship binding the tag association to a specific series from the series table.';


--
-- Name: users_following_series fk_users_following_series_series_id; Type: FK CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY users_following_series
    ADD CONSTRAINT fk_users_following_series_series_id FOREIGN KEY (series_id) REFERENCES series(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: CONSTRAINT fk_users_following_series_series_id ON users_following_series; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT fk_users_following_series_series_id ON users_following_series IS 'Foreign key binding the series identifier to an actual entry in the series table. \n\nCascade on delete and update is to remove following status from series that no longer exist.';


--
-- Name: users_following_series fk_users_following_series_user_id; Type: FK CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY users_following_series
    ADD CONSTRAINT fk_users_following_series_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: CONSTRAINT fk_users_following_series_user_id ON users_following_series; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT fk_users_following_series_user_id ON users_following_series IS 'Foreign key binding the user identifier to an actual entry in the users table. \n\nCascade on update and delete is to ensure entries for users that are removed are pruned as well.';


--
-- Name: users_groups fk_users_groups_group_name; Type: FK CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY users_groups
    ADD CONSTRAINT fk_users_groups_group_name FOREIGN KEY (group_name) REFERENCES groups_scanlation(name) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: CONSTRAINT fk_users_groups_group_name ON users_groups; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT fk_users_groups_group_name ON users_groups IS 'Foreign key binding the name of the scanlation group to an actual entry in the scanlation groups table.';


--
-- Name: users_groups fk_users_groups_users; Type: FK CONSTRAINT; Schema: public; Owner: manga
--

ALTER TABLE ONLY users_groups
    ADD CONSTRAINT fk_users_groups_users FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: CONSTRAINT fk_users_groups_users ON users_groups; Type: COMMENT; Schema: public; Owner: manga
--

COMMENT ON CONSTRAINT fk_users_groups_users ON users_groups IS 'Foreign key binding the user identifier to an actual entry in the users table.';


--
-- PostgreSQL database dump complete
--

