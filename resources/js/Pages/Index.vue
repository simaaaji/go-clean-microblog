<template>
  <div class="container">
    <h1>Microblog</h1>
    
    <!-- Post Form -->
    <div class="post-form">
      <h2>Create a Post</h2>
      <form @submit.prevent="submitPost">
        <textarea
          v-model="form.content"
          placeholder="What's on your mind?"
          rows="4"
          required
        ></textarea>
        <button type="submit">Post</button>
      </form>
    </div>

    <!-- Posts List -->
    <div class="posts-list">
      <h2>Recent Posts</h2>
      <div v-if="posts && posts.length > 0">
        <div v-for="post in posts" :key="post.id" class="post-item">
          <p>{{ post.content }}</p>
          <small>Post #{{ post.id }}</small>
        </div>
      </div>
      <div v-else>
        <p>No posts yet. Be the first to post!</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    posts: {
      type: Array,
      default: () => []
    }
  },
  data() {
    return {
      form: {
        content: ''
      }
    }
  },
  methods: {
    submitPost() {
      this.$inertia.post('/posts', this.form, {
        onSuccess: () => {
          this.form.content = ''
        }
      })
    }
  }
}
</script>

<style scoped>
.container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
  font-family: Arial, sans-serif;
}

h1 {
  color: #333;
  margin-bottom: 30px;
}

.post-form {
  background: #f5f5f5;
  padding: 20px;
  border-radius: 8px;
  margin-bottom: 30px;
}

.post-form h2 {
  margin-top: 0;
  color: #555;
}

.post-form textarea {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
  resize: vertical;
  box-sizing: border-box;
}

.post-form button {
  margin-top: 10px;
  padding: 10px 20px;
  background: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
}

.post-form button:hover {
  background: #0056b3;
}

.posts-list h2 {
  color: #555;
  margin-bottom: 20px;
}

.post-item {
  background: white;
  border: 1px solid #ddd;
  border-radius: 4px;
  padding: 15px;
  margin-bottom: 15px;
}

.post-item p {
  margin: 0 0 10px 0;
  color: #333;
  line-height: 1.5;
}

.post-item small {
  color: #999;
}
</style>
