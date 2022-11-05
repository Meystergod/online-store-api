BEGIN;

DROP TABLE IF EXISTS public.users CASCADE;
DROP TABLE IF EXISTS public.cars CASCADE;
DROP TABLE IF EXISTS public.car_model CASCADE;
DROP TABLE IF EXISTS public.car_brand CASCADE;
DROP TABLE IF EXISTS public.car_modification CASCADE;
DROP TABLE IF EXISTS public.products CASCADE;
DROP TABLE IF EXISTS public.product_category CASCADE;
DROP TABLE IF EXISTS public.product_discount CASCADE;
DROP TABLE IF EXISTS public.product_inventory CASCADE;
DROP TABLE IF EXISTS public.products_for_cars CASCADE;

END;
