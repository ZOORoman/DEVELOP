<template>
  <div>
    <h2>Todos.vue</h2>
    <router-link to="/">Home</router-link>
    <hr />
    <select v-model="filter">
      <option value="all">All</option>
      <option value="completed">Completed</option>
      <option value="not-completed">Not completed</option>
    </select>
    <hr />
    <loader v-if="loading" />
    <todolist
      v-else-if="todos.length"
      v-bind:todos="filteredTodos"
      @remove-todo="removeTodo"
    />
    <p v-else>NO todos!</p>
    <addtodo @add-todo="addTodo" />
  </div>
</template>

<script>
import todolist from "@/components/todolist";
import addtodo from "@/components/addtodo";
import loader from "@/components/loader";
export default {
  name: "app",
  data() {
    return {
      todos: [],
      //   todos: [
      //     {id: 1, title: 'Чай', completed: false},
      //     {id: 2, title: 'Хлеб', completed: false},
      //     {id: 3, title: 'Сахар', completed: false},
      //   ]
      loading: true,
      filter: "all",
    };
  },
  mounted() {
    fetch("https://jsonplaceholder.typicode.com/todos?_limit=3")
      .then((response) => response.json())
      .then((json) => {
        setTimeout(() => {
          this.todos = json;
          this.loading = false;
        }, 1000);
      });
  },
  methods: {
    removeTodo(id) {
      this.todos = this.todos.filter((t) => t.id !== id);
    },
    addTodo(todo) {
      this.todos.push(todo);
    },
  },
  // 'watch' - Позволяет сделать фильтр по жанрам, но мы воспользуемся 'computed' чтобы брать конкретные значения
  // watch: {
  //   filter(value) {
  //     console.log(value)
  //   }
  // },
  computed: {
    filteredTodos() {
      if (this.filter === 'all') {
        return this.todos;
      }
      if (this.filter === 'completed') {
        return this.todos.filter(t => t.completed);
      }
      if (this.filter === 'not-completed') {
        return this.todos.filter(t => !t.completed);
      }
    },
  },
  components: {
    todolist,
    addtodo,
    loader,
  },
};
</script>