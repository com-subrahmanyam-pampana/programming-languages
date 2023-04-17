package c_advanced.e_generics.genericmethods;

public class MyGenericMethod1 {

	public static void main(String args[]) {
		printDataType(5);
		printDataType("I am string");
		
	}
	
	private static  <T,U> void   printDataType(T myData) {
		if(myData instanceof Integer){
			System.out.println("Integer entered");
		}
		else if(myData instanceof String){
			System.out.println("String entered");
		}
		
	}
}
