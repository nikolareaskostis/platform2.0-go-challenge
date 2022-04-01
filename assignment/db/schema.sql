--Table assets

CREATE TABLE IF NOT EXISTS public.assets
(
    id uuid NOT NULL,
    description character varying COLLATE pg_catalog."default" NOT NULL,
    favourite boolean,
    "creationDate" date,
    "updateDate" date,
    CONSTRAINT assets_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.assets
    OWNER to postgres;

--Table audiences
CREATE TABLE IF NOT EXISTS public.audiences
(
    id uuid NOT NULL,
    asset_id uuid NOT NULL,
    country character varying COLLATE pg_catalog."default",
    age_min integer,
    age_max integer,
    creationdate date,
    updatedate date,
    gender gender,
    CONSTRAINT audiences_pkey PRIMARY KEY (id),
    CONSTRAINT audiences_asset_id_fkey FOREIGN KEY (asset_id)
        REFERENCES public.assets (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.audiences
    OWNER to postgres;

--Table charts
CREATE TABLE IF NOT EXISTS public.charts
(
    id uuid NOT NULL,
    asset_id uuid NOT NULL,
    x_description character varying COLLATE pg_catalog."default" NOT NULL,
    y_description character varying COLLATE pg_catalog."default" NOT NULL,
    "creationDate" date,
    "updateDate" date,
    CONSTRAINT charts_pkey PRIMARY KEY (id),
    CONSTRAINT charts_asset_id_fkey FOREIGN KEY (asset_id)
        REFERENCES public.assets (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.charts
    OWNER to postgres;

--Table chart_points
CREATE TABLE IF NOT EXISTS public.chart_points
(
    id uuid NOT NULL,
    chart_id uuid NOT NULL,
    x_val real NOT NULL,
    y_val real NOT NULL,
    CONSTRAINT chart_points_pkey PRIMARY KEY (id),
    CONSTRAINT chart_points_chart_id_fkey FOREIGN KEY (chart_id)
        REFERENCES public.charts (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.chart_points
    OWNER to postgres;

--Table insights
CREATE TABLE IF NOT EXISTS public.insights
(
    id uuid NOT NULL,
    asset_id uuid NOT NULL,
    description character varying COLLATE pg_catalog."default",
    CONSTRAINT insights_pkey PRIMARY KEY (id),
    CONSTRAINT insights_asset_id_fkey FOREIGN KEY (asset_id)
        REFERENCES public.assets (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.insights
    OWNER to postgres;

--Table user_assets
CREATE TABLE IF NOT EXISTS public.user_assets
(
    asset_id uuid NOT NULL,
    user_id uuid NOT NULL,
    "creationDate" date,
    "updateDate" date,
    CONSTRAINT user_assets_pkey PRIMARY KEY (asset_id, user_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.user_assets
    OWNER to postgres;

--Table users

CREATE TABLE IF NOT EXISTS public.users
(
    id uuid NOT NULL,
    username character varying COLLATE pg_catalog."default" NOT NULL,
    password character varying COLLATE pg_catalog."default" NOT NULL,
    email character varying COLLATE pg_catalog."default" NOT NULL,
    "creationDate" date,
    "updateDate" date,
    CONSTRAINT users_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    OWNER to postgres;