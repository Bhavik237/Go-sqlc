
CREATE SCHEMA IF NOT EXISTS course;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE IF NOT EXISTS course.course_information
(
    id serial PRIMARY KEY,
    nos_code jsonb,
    qp_code jsonb,
    sub_sector jsonb,
    sub_course jsonb,
    nsqf_level integer,
    course_id uuid,
    tags jsonb,
    sector jsonb
   
);


CREATE TABLE IF NOT EXISTS course.certificate_detail
(
    id serial PRIMARY KEY,
    certificate_enabled boolean,
    certificate_expiry_date date,
    certificate_provider_type_id character varying(255),
    certificate_provider_type integer,
    certificate_type character varying(255)
    
);

CREATE TABLE IF NOT EXISTS course.course_statistics
(
    id serial PRIMARY KEY,
    enrollment_count bigint,
    rating_average double precision,
    rating_total integer,
    course_id character varying(255)
);

CREATE TABLE IF NOT EXISTS course.course
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    code character varying(255),
    course_mode integer,
    course_status_id bigint,
    course_type integer,
    created_by character varying(255),
    created_date timestamp without time zone,
    language character varying(255),
    learning_outcome character varying(8096),
    long_description character varying(20000),
    short_description character varying(8096),
    title character varying(1024),
    updated_date timestamp without time zone,
    certificate_detail_id bigint REFERENCES course.certificate_detail(id) on delete cascade,   --realtion one to one certificate_detail
    course_information_id bigint REFERENCES course.course_information(id) on delete cascade,   --realtion one to one course_information
    course_statistics_id bigint REFERENCES course.course_statistics(id) on delete cascade,     --realtion one to one course_statistics
    end_date character varying,
    start_date character varying,
    course_provider_id character varying,
    type_id bigint,
    price double precision,
    availability bigint,
    assessment bigint,
    url character varying,
    source_system_name character varying(255),
    source_system_id character varying(255),
    learning_partner_type character varying(255),
    course_code_id  serial,
    readable_code character varying,
    sid_certificate_id character varying(1000),
    additional_amount_to_be_paid_to_partner double precision,
    lp_certificate_id character varying(255),
    is_enrollment_disabled boolean DEFAULT false,
    credit double precision,
    program_by character varying,
    schema_id character varying,
    external_payment bigint,
    is_pre_assessment boolean DEFAULT false,
    pre_assessment_type character varying,
    project_name character varying,
    CONSTRAINT course_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS course.course_occupation
(
    id serial PRIMARY KEY,
    occupation character varying(255),
    course_id uuid REFERENCES course.course(id) on delete cascade 
);
