<template>
  <ul>
    <li><router-link to='/'>Home</router-link></li>
    <li><router-link to='/orders'>Orders</router-link></li>
  </ul>

  
</template>

<script>
export default {
  name: 'HeaderComponent',
  inheritAttrs: false, // disable 'non-props' warning
};
</script>

<style>
ul {
  height: 35px;
  list-style-type: none;
  margin: 0;
  padding: 0;
  overflow: hidden;
  background-color: #333333;
}

li {
  float: left;
}

li a {
  display: block;
  color: white;
  text-align: center;
  padding: 8px;
  text-decoration: none;
}

li a:hover {
  background-color: #111111;
}
</style>
