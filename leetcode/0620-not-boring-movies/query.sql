SELECT id, movie, description, rating FROM cinema WHERE id%2=1 AND LOWER(description) != 'boring' ORDER BY rating DESC;
