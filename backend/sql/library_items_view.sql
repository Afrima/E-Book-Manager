CREATE OR REPLACE VIEW LIBRARY_ITEMS AS
SELECT
    COALESCE(COLLECTIONS.ID, BOOKS.ID) AS ID
  , COALESCE(COLLECTIONS.TITLE, BOOKS.TITLE) AS TITLE
  , CASE WHEN COLLECTIONS.TITLE IS NOT NULL THEN 'collection' ELSE 'book' END AS ITEM_TYPE
  , COUNT(BOOKS.*) AS BOOK_COUNT
  , CASE WHEN COLLECTIONS.TITLE IS NOT NULL THEN COLLECTIONS.COVER ELSE BOOKS.COVER END AS COVER
FROM
    BOOKS
        LEFT JOIN COLLECTIONS ON BOOKS.COLLECTION_ID = COLLECTIONS.ID
GROUP BY
    COALESCE(COLLECTIONS.ID, BOOKS.ID), COALESCE(COLLECTIONS.TITLE, BOOKS.TITLE), COLLECTIONS.TITLE
                                      , CASE WHEN COLLECTIONS.TITLE IS NOT NULL THEN COLLECTIONS.COVER ELSE BOOKS.COVER END;