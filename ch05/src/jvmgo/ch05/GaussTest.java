mport java.util.Scanner;

class Complex{
	    private double shi;
	        private double xu;
		    public Complex() {

			        }
		        public Complex(double s,double x) {
				        this.shi = s;
					        this.xu = x;
						    }

			    public double getShi() {
				            return shi;
					        }
			        public double getXu() {
					        return xu;
						    }
				    public void setRealPart(double s) {
					            this.shi = s;
						        }
				        public void setImaginaryPart(double x) {
						        this.xu = x;
							    }
					    public Complex add(Complex a) {
						            Complex complex = new Complex();
							            complex.shi = this.shi+a.getShi();
								            complex.xu = this.xu+a.getXu();
									            return complex;
										        }
					        public Complex sub(Complex a) {
							        Complex complex = new Complex();
								        complex.shi = this.shi-a.getShi();
									        complex.xu = this.xu-a.getXu();
										        return complex;
											    }


						    public String toString() {
							            if(shi==0&&xu==0) {
									                return 0+"";
											        }
								            if(shi==0) {
										                return xu+"i";
												        }
									            if(xu>0) {
											                return shi + "+" + xu + "i";
													        }
										            if(xu==0) {
												                return shi+"";
														        }
											            return shi +""+ xu + "i";
												        }
}

public class GaussTest
{

	    public static void main(String[] args)
		        {
				        Complex a=new Complex();
					        Complex b=new Complex();
						        Scanner in=new Scanner(System.in);
							        a.setRealPart(in.nextDouble());
								        a.setImaginaryPart(in.nextDouble());
									        b.setRealPart(in.nextDouble());
										        b.setImaginaryPart(in.nextDouble());
											        System.out.println(a);
												        System.out.println(b);
													        System.out.println(a.add(b));
														        System.out.println(a.sub(b));
															    }
}

