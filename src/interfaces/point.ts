import { prop, Typegoose, ModelType, InstanceType } from 'typegoose';

export default class IPoint {

	@prop({ required: true })
	type: string; //Point

	@prop({ required: true})
	coordinates: Array<number>;
}
