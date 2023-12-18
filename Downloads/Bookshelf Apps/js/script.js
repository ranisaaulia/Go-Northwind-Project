let books = [];
const STORAGE_KEY = 'Library_Store';

document.addEventListener('DOMContentLoaded', function () {
    const addButton = document.getElementById('addBookForm');
    addButton.addEventListener('submit', function(event) {
      event.preventDefault();
      addBook();
    });

    const searchForm = document.getElementById('searchBookForm');
    searchForm.addEventListener('submit', function(event) {
        event.preventDefault();
        searchBooks();
    });

    loadBooksFromStorage();

  });

  function renderBooks() {
    const uncompletedShelf = document.getElementById('uncompletedShelf');
    const completedShelf = document.getElementById('completedShelf');

    uncompletedShelf.innerHTML = '';
    completedShelf.innerHTML = '';

    for (const book of books) {
        const bookElement = makeBookElement(book);
        if (book.isComplete) {
            completedShelf.appendChild(bookElement);
        } else {
            uncompletedShelf.appendChild(bookElement);
        }
    }
}

function renderSearchResults(results) {
    const uncompletedShelf = document.getElementById('uncompletedShelf');
    const completedShelf = document.getElementById('completedShelf');

    uncompletedShelf.innerHTML = '';
    completedShelf.innerHTML = '';

    results.forEach(book => {
        const bookElement = makeBookElement(book);
        if (book.isComplete) {
            addBookToCompletedShelf(bookElement);
        } else {
            addBookToUncompletedShelf(bookElement);
        }
    });
}


function searchBooks() {
    const searchQuery = document.getElementById('searchTitleInput').value.toLowerCase();
  
    const searchResults = books.filter(book => {
        const title = book.title.toLowerCase();
        return title.includes(searchQuery);
    });

    renderSearchResults(searchResults);
}

function addBook() {
    const title = document.getElementById('titleInput').value;
    const author = document.getElementById('authorInput').value;
    const year = document.getElementById('yearInput').value;
    const isComplete = document.getElementById('isComplete').checked;
  
    const generatedID = generateId();
    const bookObject = generateBookObject(generatedID, title, author, parseInt(year), isComplete);
  
    books.push(bookObject);
    saveBooksToStorage();
    renderBooks();
  
    document.getElementById('titleInput').value = '';
    document.getElementById('authorInput').value = '';
    document.getElementById('yearInput').value = '';
  }

function generateId() {
    return +new Date();
  }
   
function generateBookObject(id, title, author, year, isComplete) {
    return {
        id,
        title,
        author,
        year,
        isComplete
    };
}

function addBookToUncompletedShelf(bookElement) {
    const uncompletedShelf = document.getElementById('uncompletedShelf');
    uncompletedShelf.appendChild(bookElement);
}

function addBookToCompletedShelf(bookElement) {
    const completedShelf = document.getElementById('completedShelf');
    completedShelf.appendChild(bookElement);
  }

function makeBookElement(book) {
  const bookElement = document.createElement("div");
  bookElement.classList.add('book', 'shadow');

  const bookTitle = document.createElement('h2');
  bookTitle.innerText = book.title;

  const bookAuthor = document.createElement('p');
  bookAuthor.innerText = `Author: ${book.author}`;

  const bookYear = document.createElement('p');
  bookYear.innerText = `Year: ${book.year}`;

  const bookStatus = document.createElement('p');
  bookStatus.innerText = book.isComplete ? 'Status: Selesai dibaca' : 'Status: Belum selesai dibaca';

  const deleteButton = document.createElement('button');
  deleteButton.innerText = 'Delete';
  deleteButton.addEventListener('click', function () {
      removeBook(book.id); 
      renderBooks();
  });

  const checkbox = document.createElement('input');
  checkbox.type = 'checkbox';
  checkbox.checked = book.isComplete;
  checkbox.classList.add('checkbox');

  const completionLabel = document.createElement('label');
  completionLabel.innerText = 'Completed';
  completionLabel.appendChild(checkbox);

  bookElement.appendChild(bookTitle);
  bookElement.appendChild(bookAuthor);
  bookElement.appendChild(bookYear);
  bookElement.appendChild(bookStatus);
  bookElement.appendChild(deleteButton);
  bookElement.appendChild(completionLabel);

  if (book.isComplete) {
    addBookToCompletedShelf(bookElement);
    } else {
    addBookToUncompletedShelf(bookElement);
    }

  handleBookCompletion(bookElement, book);
  
  return bookElement;
}

function handleBookCompletion(bookElement, book) {
    const checkbox = bookElement.querySelector(".checkbox");
   
    checkbox.addEventListener("change", function () {
      book.isComplete = checkbox.checked;
      saveBooksToStorage();
      renderBooks();
    });
  }

function removeBook(bookId) {
    books = books.filter(book => book.id !== bookId);
    saveBooksToStorage();
    renderBooks();
}

function saveBooksToStorage() {
  localStorage.setItem(STORAGE_KEY, JSON.stringify(books));
}

function loadBooksFromStorage() {
    const storedBooks = localStorage.getItem(STORAGE_KEY);
    if (storedBooks) {
      books = JSON.parse(storedBooks);
      renderBooks();
    }
  }