package d_collections.list;

import java.util.Stack;
public class MyStack {
	
	public static void main(String[] args) {
        Stack<String> animals= new Stack<>();
        // Add elements to Stack
        animals.push("Dog");
        animals.push("Horse");
        animals.push("Cat");
        System.out.println("Stack: " + animals);

    }
}
