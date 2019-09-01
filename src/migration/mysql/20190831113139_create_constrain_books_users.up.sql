ALTER TABLE `books` 
ADD CONSTRAINT `books_users` 
FOREIGN KEY(author_id) 
REFERENCES users(id)