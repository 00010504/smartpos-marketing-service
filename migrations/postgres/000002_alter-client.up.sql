ALTER TABLE IF EXISTS "client" ADD COLUMN "external_id" VARCHAR NOT NULL;

CREATE OR REPLACE FUNCTION create_client_external_id()
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
  AS
$$
DECLARE
  external_id INT := 0;

BEGIN
    SELECT
      COUNT(*) AS total
    FROM
      "client"
    WHERE
      "company_id" = NEW."company_id"
    INTO
      external_id;

    external_id := external_id + 1;

    NEW."external_id"=RIGHT(CONCAT('000000000', external_id), 9);

    RETURN NEW;
END;
$$;

-- triggers
CREATE OR REPLACE TRIGGER create_client_external_id
    BEFORE INSERT ON "client"
    FOR EACH ROW
    EXECUTE PROCEDURE create_client_external_id();
