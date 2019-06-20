import React from "react";
import axios from "axios";
import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import ListGroup from "react-bootstrap/ListGroup";
import Spinner from "react-bootstrap/Spinner";

export default class BookList extends React.Component {
  constructor(props) {
    super(props);

    this.state = {};
  }

  loadBook(book) {
    const location = {
      pathname: `/${book}/intro`,
      state: { fromBookList: true }
    };

    this.props.history.push(location);
  }

  render() {
    return (
      <Container>
        <Row>
          <h1>Welcome to choose your own story</h1>
        </Row>
        <Row>
          <h3>Please book from the list of books below</h3>
        </Row>

        <Row>
          {this.state.bookList ? (
            <ListGroup>
              {this.state.bookList.map((book, index) => {
                return (
                  <ListGroup.Item
                    action
                    key={index}
                    onClick={() => {
                      this.loadBook(book);
                    }}
                  >
                    {book}
                  </ListGroup.Item>
                );
              })}
            </ListGroup>
          ) : (
            <Spinner animation="border" role="status">
              <span className="sr-only">Loading...</span>
            </Spinner>
          )}
        </Row>
      </Container>
    );
  }

  componentDidMount() {
    axios.get("/booklist").then(response => {
      this.setState({ bookList: response.data });
    });
  }
}
