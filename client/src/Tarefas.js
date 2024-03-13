import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon } from "semantic-ui-react";

let endpoint = "http://localhost:9000";

class Tarefas extends Component {
    constructor(props) {
        super(props);

        this.state = {
            task: "",
            items: [],
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
        // Requisição GET para o endpoint especificado para obter a lista de tarefas
        axios.get(endpoint + "/api/task").then((res) => {
            if (res.data) { //se a resposta tiver dados então...
                // Se houver dados, atualiza o estado do componente com base nos dados recebidos
                this.setState({
                    items: res.data.map((item) => {
                        let color = "yellow";   // Define a cor inicial do cartão como amarelo
                        let style = {
                            wordWrap: "break-word" // Define o estilo do cartão
                        } 
                        if (item.status) { // Verifica o status da tarefa  e muda a cor do cartão
                            color = "green";
                            style["textDecorationLine"] = "line-through";
                        }
    
                        // Retorna um elemento de cartão para cada item de tarefa, com base nos dados recebidos
                        return (
                            <Card key={item._id} color={color} fluid>
                                <Card.Content>
                                    <Card.Header textAlign="left">
                                        <div style={style}>{item.task}</div>
                                    </Card.Header>
    
                                    {/* Botões de ação para concluir, desfazer e excluir a tarefa */}
                                    <Card.Meta textAlign="right">
                                        <Icon
                                            name="check circle"
                                            color="green"
                                            onClick={() => this.updateTask(item._id)} // Chama a função para marcar a tarefa como concluída
                                        />
                                        <span style={{ paddingRight: 10 }}>Done</span>
                                        <Icon
                                            name="undo"
                                            color="yellow"
                                            onClick={() => this.undoTask(item._id)} // Chama a função para desfazer a conclusão da tarefa
                                        />
                                        <span style={{ paddingRight: 10 }}>Undo</span>
                                        <Icon
                                            name="delete"
                                            color="red"
                                            onClick={() => this.deleteTask(item._id)} // Chama a função para excluir a tarefa
                                        />
                                        <span style={{ paddingRight: 10 }}>Delete</span>
                                    </Card.Meta>
                                </Card.Content>
                            </Card>
                        );
                    })
                })
            } else {
                this.setState({ 
                    items: [], // Se não houver dados, define o estado do componente como uma lista vazia
                });
            }
        });
    };    

    updateTask = (id) => {
        axios
            // envio de solicitacao put
            .put(endpoint + "/api/task/" + id, { // id da tarefa a ser atualizada
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded",
                },
            })
            .then((res) => {
                console.log(res);
                this.getTask();
            });
    };

    undoTask = (id) => {// Envia uma requisição PUT para desfazer uma tarefa específica no servidor
        axios
            .put(endpoint + "/api/undoTask/" + id, { // URL completa para desfazer a tarefa com o ID fornecido
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded",
                },
            })
            .then((res) => { // Ação a ser executada após a conclusão da requisição
                console.log(res);
                this.getTask(); // Após DESFAZER a tarefa com sucesso, atualiza a lista de tarefas exibida
            });
    };

    render() {
        return (
            <div>
                <div className="row">
                    <Header className="header" as="h2" color="yewllow">
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