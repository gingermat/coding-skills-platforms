SELECT email FROM (
  SELECT email, COUNT(email) AS c FROM Person
    GROUP BY email
) AS tmp WHERE c > 1;
