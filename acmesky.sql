--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2 (Debian 16.2-1.pgdg130+1)
-- Dumped by pg_dump version 16.2 (Debian 16.2-1.pgdg130+1)

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
-- Name: available_flights; Type: TABLE; Schema: public; Owner: jo
--

CREATE TABLE public.available_flights (
    id bigint NOT NULL UNIQUE,
    created_at timestamp with time zone,
    arrival_time timestamp with time zone,
    arrival_airport text,
    code text,
    cost numeric,
    offer_sent boolean,
    user_id bigint,
    airline text,
    interest_id bigint,
    departure_time timestamp with time zone,
    departure_airport text
);


ALTER TABLE public.available_flights OWNER TO jo;

--
-- Name: available_flights_id_seq; Type: SEQUENCE; Schema: public; Owner: jo
--

CREATE SEQUENCE public.available_flights_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.available_flights_id_seq OWNER TO jo;

--
-- Name: available_flights_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jo
--

ALTER SEQUENCE public.available_flights_id_seq OWNED BY public.available_flights.id;


--
-- Name: interests; Type: TABLE; Schema: public; Owner: jo
--

CREATE TABLE public.interests (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    arrival_time timestamp with time zone,
    arrival_airport text,
    user_id bigint,
    flight1_departaure_time timestamp with time zone,
    flight1_departaure_airport text,
    flight1_arrival_time timestamp with time zone,
    flight1_arrival_airport text,
    flight2_departaure_time timestamp with time zone,
    flight2_departaure_airport text,
    flight2_arrival_time timestamp with time zone,
    flight2_arrival_airport text,
    flight1_departure_time timestamp with time zone,
    flight1_departure_airport text,
    flight2_departure_time timestamp with time zone,
    flight2_departure_airport text
);


ALTER TABLE public.interests OWNER TO jo;

--
-- Name: interests_id_seq; Type: SEQUENCE; Schema: public; Owner: jo
--

CREATE SEQUENCE public.interests_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.interests_id_seq OWNER TO jo;

--
-- Name: interests_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jo
--

ALTER SEQUENCE public.interests_id_seq OWNED BY public.interests.id;


--
-- Name: journeys; Type: TABLE; Schema: public; Owner: jo
--

CREATE TABLE public.journeys (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    flight1_id bigint,
    flight2_id bigint,
    cost numeric,
    user_id bigint
);


ALTER TABLE public.journeys OWNER TO jo;

--
-- Name: journeys_id_seq; Type: SEQUENCE; Schema: public; Owner: jo
--

CREATE SEQUENCE public.journeys_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.journeys_id_seq OWNER TO jo;

--
-- Name: journeys_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jo
--

ALTER SEQUENCE public.journeys_id_seq OWNED BY public.journeys.id;


--
-- Name: offers; Type: TABLE; Schema: public; Owner: jo
--

CREATE TABLE public.offers (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    message text,
    expired text,
    token text,
    user_id bigint,
    is_used boolean,
    cost numeric,
    journey_id bigint,
    payment_link text,
    payment_paid boolean
);


ALTER TABLE public.offers OWNER TO jo;

--
-- Name: offers_id_seq; Type: SEQUENCE; Schema: public; Owner: jo
--

CREATE SEQUENCE public.offers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.offers_id_seq OWNER TO jo;

--
-- Name: offers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jo
--

ALTER SEQUENCE public.offers_id_seq OWNED BY public.offers.id;


--
-- Name: rents; Type: TABLE; Schema: public; Owner: jo
--

CREATE TABLE public.rents (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    name text,
    latitude numeric,
    longitude numeric,
    endpoint text
);


ALTER TABLE public.rents OWNER TO jo;

--
-- Name: rents_id_seq; Type: SEQUENCE; Schema: public; Owner: jo
--

CREATE SEQUENCE public.rents_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.rents_id_seq OWNER TO jo;

--
-- Name: rents_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jo
--

ALTER SEQUENCE public.rents_id_seq OWNED BY public.rents.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: jo
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    username text,
    password text,
    address text,
    name text,
    prontogram_username text,
    email text,
    is_admin boolean
);


ALTER TABLE public.users OWNER TO jo;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: jo
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO jo;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jo
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: available_flights id; Type: DEFAULT; Schema: public; Owner: jo
--

ALTER TABLE ONLY public.available_flights ALTER COLUMN id SET DEFAULT nextval('public.available_flights_id_seq'::regclass);


--
-- Name: interests id; Type: DEFAULT; Schema: public; Owner: jo
--

ALTER TABLE ONLY public.interests ALTER COLUMN id SET DEFAULT nextval('public.interests_id_seq'::regclass);


--
-- Name: journeys id; Type: DEFAULT; Schema: public; Owner: jo
--

ALTER TABLE ONLY public.journeys ALTER COLUMN id SET DEFAULT nextval('public.journeys_id_seq'::regclass);


--
-- Name: offers id; Type: DEFAULT; Schema: public; Owner: jo
--

ALTER TABLE ONLY public.offers ALTER COLUMN id SET DEFAULT nextval('public.offers_id_seq'::regclass);


--
-- Name: rents id; Type: DEFAULT; Schema: public; Owner: jo
--

ALTER TABLE ONLY public.rents ALTER COLUMN id SET DEFAULT nextval('public.rents_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: jo
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: available_flights; Type: TABLE DATA; Schema: public; Owner: jo
--

COPY public.available_flights (id, created_at, arrival_time, arrival_airport, code, cost, offer_sent, user_id, airline, interest_id, departure_time, departure_airport) FROM stdin;
39	2024-05-24 16:29:27.424701+02	2024-06-01 15:00:00+02	CTA	FR2412	350.2	t	1	http://localhost:8080/v1	36	2024-06-01 13:30:00+02	CPH
\.


--
-- Data for Name: interests; Type: TABLE DATA; Schema: public; Owner: jo
--

COPY public.interests (id, created_at, arrival_time, arrival_airport, user_id, flight1_departaure_time, flight1_departaure_airport, flight1_arrival_time, flight1_arrival_airport, flight2_departaure_time, flight2_departaure_airport, flight2_arrival_time, flight2_arrival_airport, flight1_departure_time, flight1_departure_airport, flight2_departure_time, flight2_departure_airport) FROM stdin;
40	2024-05-24 15:44:05.668981+02	\N	\N	1	\N	\N	2024-06-03 03:00:00+02	CTA	\N	\N	2024-07-13 03:00:00+02	CPH	2024-06-02 02:00:00+02	CPH	2024-07-11 02:00:00+02	CTA
37	2024-05-24 13:50:06.619294+02	\N	\N	1	\N	\N	2024-06-03 03:00:00+02	CTA	\N	\N	\N	\N	2024-06-02 02:00:00+02	CPH	\N	\N
39	2024-05-24 14:02:09.207481+02	\N	\N	6	\N	\N	2024-06-24 03:00:00+02	CTA	\N	\N	2024-06-30 03:00:00+02	CPH	2024-06-24 02:00:00+02	CPH	2024-06-30 02:00:00+02	BLQ
36	2024-05-24 13:42:07.298503+02	2024-06-01 03:00:00+02	\N	1	\N	\N	2024-06-01 13:00:00+02	CTA	\N	\N	\N	\N	2024-06-01 10:00:00+02	CPH	\N	\N
\.


--
-- Data for Name: journeys; Type: TABLE DATA; Schema: public; Owner: jo
--

COPY public.journeys (id, created_at, flight1_id, flight2_id, cost, user_id) FROM stdin;
55	2024-05-24 17:41:14.879505+02	39	\N	350.2	1
\.


--
-- Data for Name: offers; Type: TABLE DATA; Schema: public; Owner: jo
--

COPY public.offers (id, created_at, message, expired, token, user_id, is_used, cost, journey_id, payment_link, payment_paid) FROM stdin;
81	2024-05-24 17:41:14.917591+02	Hello Mario Rossi, this is the offer token for your flight from <b>CPH</b> to <b>CTA</b> in date 01/06/2024 13:30 - 01/06/2024 15:00 for 350.20€. <br>The total for your journey is 350.20€. <br><a href="#" target="_blank">UZFWYT</a>	1716651674	UZFWYT	1	t	\N	55	http://localhost:5173/?id=99d378c6-da03-4b4a-aaf2-c2fdb3c3f580	f
\.


--
-- Data for Name: rents; Type: TABLE DATA; Schema: public; Owner: jo
--

COPY public.rents (id, created_at, name, latitude, longitude, endpoint) FROM stdin;
1	\N	Uber	37.512157	15.0802861	https://gist.githubusercontent.com/boozec/5da4702c0f306b83b836eb3039c678a4/raw/e10f3f4f24dba344eadf2b5a3458d73c045377e3/rent.wsdl
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: jo
--

COPY public.users (id, created_at, updated_at, deleted_at, username, password, address, name, prontogram_username, email, is_admin) FROM stdin;
4	2024-05-23 15:59:19.159649+02	\N	\N	sa3	ciaone123	Mura Anteo Zamboni 8, Bologna	Mario Draghi	\N	mario@u2	f
5	2024-05-23 15:59:43.285839+02	\N	\N	sa3	ciaone123	Mura Anteo Zamboni 8, Bologna	Mario Draghi	\N	mario@unibo.it	f
6	2024-05-23 16:03:37.014106+02	\N	\N	sa4	fa4cf43b69bcd39272557958d798cd74530125d5af63d2ef9a367aa1fcb5e804	Mura Anteo Zamboni 8, Bologna	Mario Draghi	\N	mario2@unibo.it	f
7	2024-05-23 16:34:05.471027+02	\N	\N	sa5	\N	Mura Anteo Zamboni 8, Bologna	Mario Draghi	\N	mario3@unibo.it	f
8	2024-05-23 16:34:35.184481+02	\N	\N	sa6	fa4cf43b69bcd39272557958d798cd74530125d5af63d2ef9a367aa1fcb5e804	Mura Anteo Zamboni 8, Bologna	Mario Draghi	\N	mario4@unibo.it	f
9	2024-05-23 16:35:11.31336+02	\N	\N	sa7	fa4cf43b69bcd39272557958d798cd74530125d5af63d2ef9a367aa1fcb5e804	Mura Anteo Zamboni 8, Bologna	Mario Draghi	\N	mario5@unibo.it	f
2	2024-05-23 15:42:40.051386+02	\N	\N	sa	ciaone123	\N	Mario Draghi	\N	mario@u	t
3	2024-05-23 15:58:40.784977+02	\N	\N	sa	ciaone123	\N	Mario Draghi	\N	mario@u	t
1	\N	\N	\N	sa	6ea044c786f237c955b497b04b9247f2a663c5038e54175e62308c8b8457e23e	Viale Andrea Doria 6, Catania	Mario Rossi	\N	santo@unibo.it	t
\.


--
-- Name: available_flights_id_seq; Type: SEQUENCE SET; Schema: public; Owner: jo
--

SELECT pg_catalog.setval('public.available_flights_id_seq', 39, true);


--
-- Name: interests_id_seq; Type: SEQUENCE SET; Schema: public; Owner: jo
--

SELECT pg_catalog.setval('public.interests_id_seq', 40, true);


--
-- Name: journeys_id_seq; Type: SEQUENCE SET; Schema: public; Owner: jo
--

SELECT pg_catalog.setval('public.journeys_id_seq', 55, true);


--
-- Name: offers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: jo
--

SELECT pg_catalog.setval('public.offers_id_seq', 81, true);


--
-- Name: rents_id_seq; Type: SEQUENCE SET; Schema: public; Owner: jo
--

SELECT pg_catalog.setval('public.rents_id_seq', 1, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: jo
--

SELECT pg_catalog.setval('public.users_id_seq', 9, true);


--
-- Name: available_flights available_flights_pkey; Type: CONSTRAINT; Schema: public; Owner: jo
--

ALTER TABLE ONLY public.available_flights
    ADD CONSTRAINT available_flights_pkey PRIMARY KEY (id);


--
-- Name: interests interests_pkey; Type: CONSTRAINT; Schema: public; Owner: jo
--

ALTER TABLE ONLY public.interests
    ADD CONSTRAINT interests_pkey PRIMARY KEY (id);


--
-- Name: journeys journeys_pkey; Type: CONSTRAINT; Schema: public; Owner: jo
--

ALTER TABLE ONLY public.journeys
    ADD CONSTRAINT journeys_pkey PRIMARY KEY (id);


--
-- Name: offers offers_pkey; Type: CONSTRAINT; Schema: public; Owner: jo
--

ALTER TABLE ONLY public.offers
    ADD CONSTRAINT offers_pkey PRIMARY KEY (id);


--
-- Name: rents rents_pkey; Type: CONSTRAINT; Schema: public; Owner: jo
--

ALTER TABLE ONLY public.rents
    ADD CONSTRAINT rents_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: jo
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: jo
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: available_flights fk_available_flights_interest; Type: FK CONSTRAINT; Schema: public; Owner: jo
--

ALTER TABLE ONLY public.available_flights
    ADD CONSTRAINT fk_available_flights_interest FOREIGN KEY (interest_id) REFERENCES public.interests(id);


--
-- Name: available_flights fk_available_flights_user; Type: FK CONSTRAINT; Schema: public; Owner: jo
--

ALTER TABLE ONLY public.available_flights
    ADD CONSTRAINT fk_available_flights_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: interests fk_interests_user; Type: FK CONSTRAINT; Schema: public; Owner: jo
--

ALTER TABLE ONLY public.interests
    ADD CONSTRAINT fk_interests_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: journeys fk_journeys_flight1; Type: FK CONSTRAINT; Schema: public; Owner: jo
--

ALTER TABLE ONLY public.journeys
    ADD CONSTRAINT fk_journeys_flight1 FOREIGN KEY (flight1_id) REFERENCES public.available_flights(id);


--
-- Name: journeys fk_journeys_flight2; Type: FK CONSTRAINT; Schema: public; Owner: jo
--

ALTER TABLE ONLY public.journeys
    ADD CONSTRAINT fk_journeys_flight2 FOREIGN KEY (flight2_id) REFERENCES public.available_flights(id);


--
-- Name: journeys fk_journeys_user; Type: FK CONSTRAINT; Schema: public; Owner: jo
--

ALTER TABLE ONLY public.journeys
    ADD CONSTRAINT fk_journeys_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: offers fk_offers_journey; Type: FK CONSTRAINT; Schema: public; Owner: jo
--

ALTER TABLE ONLY public.offers
    ADD CONSTRAINT fk_offers_journey FOREIGN KEY (journey_id) REFERENCES public.journeys(id);


--
-- Name: offers fk_offers_user; Type: FK CONSTRAINT; Schema: public; Owner: jo
--

ALTER TABLE ONLY public.offers
    ADD CONSTRAINT fk_offers_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

