import React, {Component} from "react";
import axios from "axios";
import {Card, Header, Form, Input, Icon} from "semantic-ui-react";

let endpoint = "http://localhost:9000";

class Tarefas extends Component {
    constructor(props) {
        super(props);

        this.state ={
            task:"",
            items:[],
        };
    }
    componentDidMount() {
        this.getTasks();
    }
    onChange = (event) => {
        this.setState({
            [event.target.name]: event.target.value,
        })
    }

    getTask = () => {
        axios.get(endpoint + "/api/task").then((res)=>{
            if (res.data) {
                this.setState({
                    items: res.data.map((item) => {
                        let color = "yellow";
                        let style = {
                            wordWrap: "break-word"
                        }

                        if (item.status) {
                            color = "green";
                            style["textDecorationLine"] = "line-through";
                        }

                        return (
                            <Card key={item._id} color={color} fluid>
                              <Card.Content>
                                <Card.Header textAlign="left">
                                  <div style={style}>{item.task}</div>
                                </Card.Header>
              
                                <Card.Meta textAlign="right">
                                  <Icon
                                    name="check circle"
                                    color="green"
                                    onClick={() => this.updateTask(item._id)}
                                  />
                                  <span style={{ paddingRight: 10 }}>Done</span>
                                  <Icon
                                    name="undo"
                                    color="yellow"
                                    onClick={() => this.undoTask(item._id)}
                                  />
                                  <span style={{ paddingRight: 10 }}>Undo</span>
                                  <Icon
                                    name="delete"
                                    color="red"
                                    onClick={() => this.deleteTask(item._id)}
                                  />
                                  <span style={{ paddingRight: 10 }}>Delete</span>
                                </Card.Meta>
                              </Card.Content>
                            </Card>
                          );
                    })
                })
            }
        })
    }

    render() {
        return (
            <div>
                <div className="row">
                 <Header className = "header" as="h2" color="yewllow">
                    Tarefas
                 </Header>
                </div>   
                <div className="row">
                    <Form onSubmit={this.onChange}
                    value={this.state.task}
                    fluid
                    placeholder="Create Task"
                    >
                    {/* Botão Criar tarefa */}
                    </Form>
                </div> 
                <div className="row">
                    <Card.Group>
                        {this.state.items}
                    </Card.Group>
                </div>
            </div>
        );
    }
} 

export default Tarefas;