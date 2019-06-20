import React, { Component } from "react";
import axios from "axios";
import Card from "react-bootstrap/Card";
import ListGroup from "react-bootstrap/ListGroup";
import Container from "react-bootstrap/Container";
import Spinner from "react-bootstrap/Spinner";
import Breadcrumb from "react-bootstrap/Breadcrumb";

export default class StoryArc extends Component {
  constructor(props) {
    super(props);

    this.state = {
      selectedStoryArc: undefined
    };
  }

  getStoryArc(book, story) {
    axios.get(`/${book}/${story}`).then(response => {
      this.setState({ data: response.data });
    });
  }

  selectArc(story) {
    this.setState({ selectedStoryArc: story });
  }

  render() {
    return this.state.data ? (
      <Container>
        <Breadcrumb>
          <Breadcrumb.Item
            onClick={() => {
              this.props.history.replace("/");
            }}
          >
            Home
          </Breadcrumb.Item>
        </Breadcrumb>
        <Card border={"primary"}>
          <Card.Header>{this.state.data.title}</Card.Header>
          <Card.Body>
            <Card.Text>{this.state.data.story}</Card.Text>
          </Card.Body>
          <Card.Footer>
            {this.state.data.options && this.state.data.options.length > 0 ? (
              <Card>
                <Card.Title>What will you do?</Card.Title>
                <ListGroup>
                  {this.state.data.options.map((option, index) => {
                    return (
                      <ListGroup.Item
                        key={`options-${index}`}
                        action
                        onClick={() => {
                          this.selectArc(`${option.arc}`);
                        }}
                      >
                        {option.text}
                      </ListGroup.Item>
                    );
                  })}
                </ListGroup>
              </Card>
            ) : (
              <br />
            )}
          </Card.Footer>
        </Card>
      </Container>
    ) : (
      <Spinner animation="border" role="status">
        <span className="sr-only">Loading...</span>
      </Spinner>
    );
  }

  componentDidMount() {
    this.getStoryArc(
      this.props.match.params.book,
      this.props.match.params.story
    );
  }

  componentDidUpdate(prevProps, prevState, snapshot) {
    if (this.state.selectedStoryArc !== prevState.selectedStoryArc) {
      this.props.history.push(
        `/${this.props.match.params.book}/${this.state.selectedStoryArc}`
      );
      this.getStoryArc(
        this.props.match.params.book,
        this.state.selectedStoryArc
      );
    }
  }
}
