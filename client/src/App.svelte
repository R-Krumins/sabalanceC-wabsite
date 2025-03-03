<script lang="ts">
	import { onMount } from 'svelte';

	type Product = {
		name: string;
		desc: string;
		price: Number;
		img: string;
	};

	let products: Product[] = [];

	onMount(async () => {
		const res = await fetch('api/product');
		const data = await res.json();
		console.log(data);
		products = data;
	});
</script>

<main>
	<h1>Products</h1>
	{#each products as product}
		<div class="max-w-sm rounded overflow-hidden shadow-lg m-4">
			<img class="w-full" src={product.img} alt={product.name} />
			<div class="px-6 py-4">
				<h2 class="font-bold text-xl mb-2">{product.name}</h2>
				<p class="text-gray-700 text-base">{product.desc}</p>
				<p class="text-gray-900 text-lg font-semibold mt-2">${product.price}</p>
			</div>
		</div>
	{/each}
</main>
