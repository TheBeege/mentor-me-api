--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

SET search_path = public, pg_catalog;

ALTER TABLE ONLY public.mentor_topic DROP CONSTRAINT mentor_topic_user_id;
ALTER TABLE ONLY public.mentor_topic DROP CONSTRAINT mentor_topic_topic_id;
ALTER TABLE ONLY public.mentor_request DROP CONSTRAINT mentor_request_mentor_id;
ALTER TABLE ONLY public.mentor_request DROP CONSTRAINT mentor_request_mentee_id;
DROP INDEX public.fki_mentor_request_mentor_id;
DROP INDEX public.fki_mentor_request_mentee_id;
ALTER TABLE ONLY public."user" DROP CONSTRAINT user_unq_username;
ALTER TABLE ONLY public."user" DROP CONSTRAINT user_unq_email;
ALTER TABLE ONLY public."user" DROP CONSTRAINT user_pk;
ALTER TABLE ONLY public.topic DROP CONSTRAINT topic_unq_name;
ALTER TABLE ONLY public.topic DROP CONSTRAINT topic_id;
ALTER TABLE ONLY public.mentor_topic DROP CONSTRAINT mentor_topic_id;
ALTER TABLE ONLY public.mentor_request DROP CONSTRAINT mentor_request_id;
ALTER TABLE public."user" ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.topic ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.mentor_topic ALTER COLUMN mentor_topic_id DROP DEFAULT;
ALTER TABLE public.mentor_request ALTER COLUMN mentor_request_id DROP DEFAULT;
DROP SEQUENCE public.user_user_id_seq;
DROP TABLE public."user";
DROP SEQUENCE public.topic_id_seq;
DROP TABLE public.topic;
DROP SEQUENCE public.mentor_topic_mentor_topic_id_seq;
DROP TABLE public.mentor_topic;
DROP SEQUENCE public.mentor_request_mentor_request_id_seq;
DROP TABLE public.mentor_request;
DROP EXTENSION plpgsql;
DROP SCHEMA public;
--
-- Name: public; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA public;


--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON SCHEMA public IS 'standard public schema';


--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: mentor_request; Type: TABLE; Schema: public; Owner: -; Tablespace: 
--

CREATE TABLE mentor_request (
    mentor_request_id integer NOT NULL,
    mentor_id integer NOT NULL,
    mentee_id integer NOT NULL,
    requested timestamp without time zone NOT NULL,
    accepted timestamp without time zone,
    rejected timestamp without time zone
);


--
-- Name: mentor_request_mentor_request_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE mentor_request_mentor_request_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: mentor_request_mentor_request_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE mentor_request_mentor_request_id_seq OWNED BY mentor_request.mentor_request_id;


--
-- Name: mentor_topic; Type: TABLE; Schema: public; Owner: -; Tablespace: 
--

CREATE TABLE mentor_topic (
    mentor_topic_id integer NOT NULL,
    user_id integer NOT NULL,
    topic_id integer NOT NULL,
    level integer,
    description text
);


--
-- Name: mentor_topic_mentor_topic_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE mentor_topic_mentor_topic_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: mentor_topic_mentor_topic_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE mentor_topic_mentor_topic_id_seq OWNED BY mentor_topic.mentor_topic_id;


--
-- Name: topic; Type: TABLE; Schema: public; Owner: -; Tablespace: 
--

CREATE TABLE topic (
    id integer NOT NULL,
    name text NOT NULL
);


--
-- Name: topic_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE topic_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: topic_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE topic_id_seq OWNED BY topic.id;


--
-- Name: user; Type: TABLE; Schema: public; Owner: -; Tablespace: 
--

CREATE TABLE "user" (
    id integer NOT NULL,
    username text,
    display_name text,
    email text,
    created timestamp without time zone DEFAULT now(),
    last_activity timestamp without time zone DEFAULT now() NOT NULL,
    password text,
    description text,
    icon_url text
);


--
-- Name: user_user_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE user_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE user_user_id_seq OWNED BY "user".id;


--
-- Name: mentor_request_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY mentor_request ALTER COLUMN mentor_request_id SET DEFAULT nextval('mentor_request_mentor_request_id_seq'::regclass);


--
-- Name: mentor_topic_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY mentor_topic ALTER COLUMN mentor_topic_id SET DEFAULT nextval('mentor_topic_mentor_topic_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY topic ALTER COLUMN id SET DEFAULT nextval('topic_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY "user" ALTER COLUMN id SET DEFAULT nextval('user_user_id_seq'::regclass);


--
-- Name: mentor_request_id; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY mentor_request
    ADD CONSTRAINT mentor_request_id PRIMARY KEY (mentor_request_id);


--
-- Name: mentor_topic_id; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY mentor_topic
    ADD CONSTRAINT mentor_topic_id PRIMARY KEY (mentor_topic_id);


--
-- Name: topic_id; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY topic
    ADD CONSTRAINT topic_id PRIMARY KEY (id);


--
-- Name: topic_unq_name; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY topic
    ADD CONSTRAINT topic_unq_name UNIQUE (name);


--
-- Name: user_pk; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_pk PRIMARY KEY (id);


--
-- Name: user_unq_email; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_unq_email UNIQUE (email);


--
-- Name: user_unq_username; Type: CONSTRAINT; Schema: public; Owner: -; Tablespace: 
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_unq_username UNIQUE (username);


--
-- Name: fki_mentor_request_mentee_id; Type: INDEX; Schema: public; Owner: -; Tablespace: 
--

CREATE INDEX fki_mentor_request_mentee_id ON mentor_request USING btree (mentee_id);


--
-- Name: fki_mentor_request_mentor_id; Type: INDEX; Schema: public; Owner: -; Tablespace: 
--

CREATE INDEX fki_mentor_request_mentor_id ON mentor_request USING btree (mentor_id);


--
-- Name: mentor_request_mentee_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY mentor_request
    ADD CONSTRAINT mentor_request_mentee_id FOREIGN KEY (mentee_id) REFERENCES "user"(id);


--
-- Name: mentor_request_mentor_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY mentor_request
    ADD CONSTRAINT mentor_request_mentor_id FOREIGN KEY (mentor_id) REFERENCES "user"(id);


--
-- Name: mentor_topic_topic_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY mentor_topic
    ADD CONSTRAINT mentor_topic_topic_id FOREIGN KEY (topic_id) REFERENCES topic(id);


--
-- Name: mentor_topic_user_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY mentor_topic
    ADD CONSTRAINT mentor_topic_user_id FOREIGN KEY (user_id) REFERENCES "user"(id);


--
-- Name: public; Type: ACL; Schema: -; Owner: -
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

