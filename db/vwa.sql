--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: komentar; Type: TABLE; Schema: public; Owner: vwa; Tablespace: 
--

CREATE TABLE komentar (
    id_komentar integer NOT NULL,
    isi_komentar text NOT NULL,
    id_user integer NOT NULL
);


ALTER TABLE public.komentar OWNER TO vwa;

--
-- Name: komentar_id_komentar_seq; Type: SEQUENCE; Schema: public; Owner: vwa
--

CREATE SEQUENCE komentar_id_komentar_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.komentar_id_komentar_seq OWNER TO vwa;

--
-- Name: komentar_id_komentar_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: vwa
--

ALTER SEQUENCE komentar_id_komentar_seq OWNED BY komentar.id_komentar;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE users (
    id integer NOT NULL,
    username character varying(40) NOT NULL,
    password character varying(50) NOT NULL,
    email character varying(50) NOT NULL,
    phone_number character varying(12) NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE users_id_seq OWNED BY users.id;


--
-- Name: id_komentar; Type: DEFAULT; Schema: public; Owner: vwa
--

ALTER TABLE ONLY komentar ALTER COLUMN id_komentar SET DEFAULT nextval('komentar_id_komentar_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY users ALTER COLUMN id SET DEFAULT nextval('users_id_seq'::regclass);


--
-- Data for Name: komentar; Type: TABLE DATA; Schema: public; Owner: vwa
--

COPY komentar (id_komentar, isi_komentar, id_user) FROM stdin;
1	ready gan?? ini ROMnya china??	2
2	ready, bisa gosend hari ini	3
\.


--
-- Name: komentar_id_komentar_seq; Type: SEQUENCE SET; Schema: public; Owner: vwa
--

SELECT pg_catalog.setval('komentar_id_komentar_seq', 8, true);


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY users (id, username, password, email, phone_number) FROM stdin;
2	eko susilo	a119e534072584a0ea88cdea4664aecd	pakeko@gmail.com	08564771185
3	rahmat	5fe43373c2db4deb851f3290080621f5	rahmat01@gmail.com	082342844322
1	andii	6b7330782b2feb4924020cc4a57782a9	andi03@gmail.com	08882228854
\.


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('users_id_seq', 3, true);


--
-- Name: komentar_pkey; Type: CONSTRAINT; Schema: public; Owner: vwa; Tablespace: 
--

ALTER TABLE ONLY komentar
    ADD CONSTRAINT komentar_pkey PRIMARY KEY (id_komentar);


--
-- Name: users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- Name: users; Type: ACL; Schema: public; Owner: postgres
--

REVOKE ALL ON TABLE users FROM PUBLIC;
REVOKE ALL ON TABLE users FROM postgres;
GRANT ALL ON TABLE users TO postgres;
GRANT ALL ON TABLE users TO vwa;


--
-- PostgreSQL database dump complete
--

