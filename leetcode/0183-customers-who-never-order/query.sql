SELECT c.name AS Customers FROM Customers c
  WHERE id NOT IN (
    SELECT customerId FROM Orders GROUP BY customerId
  );
