--
-- PostgreSQL database dump
--

-- Dumped from database version 14.11 (Ubuntu 14.11-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.11 (Ubuntu 14.11-0ubuntu0.22.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: aircraft; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aircraft (
    id integer NOT NULL,
    serial_number character varying(50) NOT NULL,
    manufacturer character varying(100) NOT NULL
);


ALTER TABLE public.aircraft OWNER TO postgres;

--
-- Name: aircraft_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.aircraft_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.aircraft_id_seq OWNER TO postgres;

--
-- Name: aircraft_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.aircraft_id_seq OWNED BY public.aircraft.id;


--
-- Name: flight; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.flight (
    id integer NOT NULL,
    departure_airport character varying(4) NOT NULL,
    arrival_airport character varying(4) NOT NULL,
    departure_datetime timestamp without time zone NOT NULL,
    arrival_datetime timestamp without time zone NOT NULL,
    aircraft_id integer
);


ALTER TABLE public.flight OWNER TO postgres;

--
-- Name: flight_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.flight_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.flight_id_seq OWNER TO postgres;

--
-- Name: flight_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.flight_id_seq OWNED BY public.flight.id;


--
-- Name: aircraft id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aircraft ALTER COLUMN id SET DEFAULT nextval('public.aircraft_id_seq'::regclass);


--
-- Name: flight id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.flight ALTER COLUMN id SET DEFAULT nextval('public.flight_id_seq'::regclass);


--
-- Data for Name: aircraft; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aircraft (id, serial_number, manufacturer) FROM stdin;
1	ABC123	Boeing
2	DEF456	Airbus
3	PQR678	Cessna
4	MNO345	Bombardier
5	JKL012	Embraer
\.


--
-- Data for Name: flight; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.flight (id, departure_airport, arrival_airport, departure_datetime, arrival_datetime, aircraft_id) FROM stdin;
1	LFPG	EDDF	2024-05-06 08:00:00	2024-05-06 10:00:00	3
2	LEBL	LFPG	2024-05-07 08:00:00	2024-05-07 10:00:00	1
3	LEBL	LFPG	2024-05-08 08:00:00	2024-05-08 10:00:00	\N
4	LFPG	EDDF	2024-05-09 08:00:00	2024-05-09 10:00:00	\N
5	LEBL	LFPG	2024-05-10 08:00:00	2024-05-10 10:00:00	\N
6	LFPG	EDDF	2024-05-05 08:00:00	2024-05-05 10:00:00	\N
7	LEBL	LFPG	2024-05-04 08:00:00	2024-05-04 10:00:00	\N
8	LFPG	EDDF	2024-05-03 08:00:00	2024-05-03 10:00:00	2
9	LEBL	LFPG	2024-05-02 08:00:00	2024-05-02 10:00:00	1
10	LEBL	LFPG	2024-05-13 08:00:00	2024-05-13 10:00:00	\N
11	LFPG	EDDF	2024-05-14 08:00:00	2024-05-14 10:00:00	\N
12	LEBL	LFPG	2024-05-15 08:00:00	2024-05-15 10:00:00	\N
13	JFK	LHR	2024-05-16 08:00:00	2024-05-16 12:00:00	1
14	DEL	SFO	2024-05-17 10:00:00	2024-05-17 15:00:00	2
15	LAX	CDG	2024-05-18 08:00:00	2024-05-18 13:00:00	\N
17	ATL	AMS	2024-05-20 08:00:00	2024-05-20 12:00:00	3
18	MEX	FCO	2024-05-21 10:00:00	2024-05-21 15:00:00	1
19	FRA	DXB	2024-05-22 08:00:00	2024-05-22 13:00:00	\N
20	BKK	JNB	2024-05-23 09:00:00	2024-05-23 14:00:00	\N
21	YYZ	SYD	2024-05-27 09:00:00	2024-05-27 23:00:00	\N
22	IST	NRT	2024-05-26 08:00:00	2024-05-26 18:00:00	\N
23	SIN	EZE	2024-05-25 10:00:00	2024-05-25 20:00:00	3
24	LHR	ICN	2024-05-24 08:00:00	2024-05-24 15:00:00	2
16	PEK	BOM	2024-05-19 09:00:00	2024-05-19 14:00:00	5
\.


--
-- Name: aircraft_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.aircraft_id_seq', 5, true);


--
-- Name: flight_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.flight_id_seq', 24, true);


--
-- Name: aircraft aircraft_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aircraft
    ADD CONSTRAINT aircraft_pkey PRIMARY KEY (id);


--
-- Name: flight flight_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.flight
    ADD CONSTRAINT flight_pkey PRIMARY KEY (id);


--
-- Name: flight flight_aircraft_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.flight
    ADD CONSTRAINT flight_aircraft_id_fkey FOREIGN KEY (aircraft_id) REFERENCES public.aircraft(id);


--
-- Name: TABLE aircraft; Type: ACL; Schema: public; Owner: postgres
--

GRANT ALL ON TABLE public.aircraft TO fleet_manager;


--
-- Name: TABLE flight; Type: ACL; Schema: public; Owner: postgres
--

GRANT ALL ON TABLE public.flight TO fleet_manager;


--
-- PostgreSQL database dump complete
--

