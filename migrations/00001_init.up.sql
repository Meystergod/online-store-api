BEGIN;

SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = ON;
SET check_function_bodies = FALSE;
SET client_min_messages = WARNING;
SET search_path = public, extensions;
SET default_tablespace = '';
SET default_with_oids = FALSE;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE public.user
(
    id SERIAL PRIMARY KEY,
    password TEXT NOT NULL,
    first_name TEXT,
    last_name TEXT,
    created_at TIMESTAMP,
    visited_at TIMESTAMP,
    is_admin BOOLEAN NOT NULL
);

CREATE TABLE public.car_brand
(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL
);

CREATE TABLE public.car
(
    id SERIAL PRIMARY KEY,
    description TEXT NOT NULL,
    brand_id INT NOT NULL REFERENCES public.car_brand(id),
    model TEXT NOT NULL,
    modification TEXT NOT NULL
);

CREATE TABLE public.product_category
(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE public.product_discount
(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE public.product_inventory
(
    id SERIAL PRIMARY KEY,
    quantity INT
);

CREATE TABLE public.product
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    brand TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    price BIGINT NOT NULL,
    image_id UUID,
    specifications JSONB,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    category_id INT NOT NULL REFERENCES public.product_category(id),
    discount_id INT NOT NULL REFERENCES public.product_discount(id),
    inventory_id INT NOT NULL UNIQUE REFERENCES public.product_inventory(id),
    CONSTRAINT positive_price CHECK (price > 0)
);

CREATE TABLE products_for_cars
(
    product_id UUID NOT NULL REFERENCES public.product(id),
    car_id INT NOT NULL REFERENCES public.car(id),
    CONSTRAINT product_car_pkey PRIMARY KEY (product_id, car_id)
);

COMMIT;
