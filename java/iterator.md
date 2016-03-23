##遍历
ArrayList<String> list = new ArrayList<String>();
list.add("1");
list.add("2");
list.add("3");
Iterator it = list.iterator();
while(it.hasNext()){
	System.out.println(it.next());
}
输出 1 2 3