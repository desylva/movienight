--
-- PostgreSQL database dump
--

-- Dumped from database version 11.2
-- Dumped by pg_dump version 11.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: movies; Type: TABLE; Schema: public; Owner: olivier
--

CREATE TABLE public.movies (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    user_uuid character varying(255) NOT NULL,
    imdb_id character varying(255) DEFAULT ''::character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    users_for character varying[],
    users_against character varying[],
    score integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.movies OWNER TO olivier;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: olivier
--

CREATE TABLE public.schema_migration (
    version character varying(255) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO olivier;

--
-- Name: users; Type: TABLE; Schema: public; Owner: olivier
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    active boolean NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    color character varying(255),
    password_hash character varying(255) DEFAULT ''::character varying NOT NULL,
    password_hash_reset character varying(255) DEFAULT ''::character varying NOT NULL,
    verified boolean DEFAULT false NOT NULL,
    verification_hash character varying(255) DEFAULT ''::character varying NOT NULL
);


ALTER TABLE public.users OWNER TO olivier;

--
-- Name: movies movies_pkey; Type: CONSTRAINT; Schema: public; Owner: olivier
--

ALTER TABLE ONLY public.movies
    ADD CONSTRAINT movies_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: olivier
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: olivier
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: users_email_idx; Type: INDEX; Schema: public; Owner: olivier
--

CREATE UNIQUE INDEX users_email_idx ON public.users USING btree (email);


--
-- PostgreSQL database dump complete
--

