--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.5
-- Dumped by pg_dump version 9.6.5

-- Started on 2019-06-27 13:23:57

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET search_path = public, pg_catalog;

--
-- TOC entry 188 (class 1255 OID 28258)
-- Name: timestamper(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION timestamper() RETURNS trigger
    LANGUAGE plpgsql
    AS $$

begin

	CASE TG_OP

	WHEN 'INSERT' THEN

        NEW.created_on := now() at time zone 'Asia/Jakarta';

        RETURN NEW;

    WHEN 'UPDATE' THEN

		NEW.updated_on := now() at time zone 'Asia/Jakarta';

        RETURN NEW;

    END CASE;

END;

$$;


SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 187 (class 1259 OID 39813)
-- Name: sample; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE sample (
    id integer NOT NULL,
    field_a character varying NOT NULL,
    field_b double precision NOT NULL,
    field_c boolean NOT NULL,
    created_on timestamp without time zone,
    updated_on timestamp without time zone
);


--
-- TOC entry 186 (class 1259 OID 39811)
-- Name: sample_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE sample_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 2132 (class 0 OID 0)
-- Dependencies: 186
-- Name: sample_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE sample_id_seq OWNED BY sample.id;


--
-- TOC entry 2004 (class 2604 OID 39816)
-- Name: sample id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY sample ALTER COLUMN id SET DEFAULT nextval('sample_id_seq'::regclass);


--
-- TOC entry 2126 (class 0 OID 39813)
-- Dependencies: 187
-- Data for Name: sample; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO sample VALUES (1, 'abc', 0.10000000000000001, true, '2019-06-27 13:23:34.083729', NULL);


--
-- TOC entry 2133 (class 0 OID 0)
-- Dependencies: 186
-- Name: sample_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('sample_id_seq', 1, true);


--
-- TOC entry 2006 (class 2606 OID 39821)
-- Name: sample sample_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY sample
    ADD CONSTRAINT sample_pk PRIMARY KEY (id);


--
-- TOC entry 2007 (class 2620 OID 39822)
-- Name: sample timestamper; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER timestamper BEFORE INSERT OR UPDATE ON sample FOR EACH ROW EXECUTE PROCEDURE timestamper();


-- Completed on 2019-06-27 13:23:57

--
-- PostgreSQL database dump complete
--

